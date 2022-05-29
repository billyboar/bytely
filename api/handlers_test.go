//go:build integration

package api

import (
	"context"
	"sync"
	"testing"

	"github.com/billyboar/bytely/api/config"
	"github.com/billyboar/bytely/internal/storage"
	"github.com/billyboar/bytely/internal/storage/pg"
	"github.com/billyboar/bytely/pb"
	"github.com/billyboar/bytely/schema"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const defaultDbConnectionStr = "postgresql://postgres:postgres@localhost:5432/bytely?sslmode=disable"

type HandlersTestSuite struct {
	suite.Suite

	db  storage.Storage
	srv *Server
}

func (suite *HandlersTestSuite) SetupSuite() {
	var err error
	suite.db, err = pg.NewClient(defaultDbConnectionStr)
	if err != nil {
		suite.FailNow("cannot connect to the database")
	}

	suite.srv = NewServer(suite.db, &config.Config{
		Port:                  8080,
		LogLevel:              "debug",
		DatabaseConnectionStr: defaultDbConnectionStr,
	})

	suite.prepareDatabase()
}

func (suite *HandlersTestSuite) Test_addURL_OK() {
	resp, err := suite.srv.AddURL(context.Background(), &pb.AddURLRequest{
		OriginalUrl: "https://google.com",
	})

	suite.Nil(err, "error should be nil")
	suite.EqualValues(len(resp.ShortUrl), defaultURLLength, "short url should be empty")
}

func (suite *HandlersTestSuite) Test_addURL_InvalidShortURL() {
	resp, err := suite.srv.AddURL(context.Background(), &pb.AddURLRequest{
		OriginalUrl: "invalidurl",
	})

	suite.Nil(resp, "response should be nil")
	suite.Require().NotNil(err, "error should not be nil")

	if errStatus, ok := status.FromError(err); ok {
		suite.Equal(errStatus.Code(), codes.InvalidArgument, "error code should be invalid argument (3)")
	} else {
		suite.Fail("error should be a GRPC status error")
	}
}

func (suite *HandlersTestSuite) Test_getOriginalURL_OK() {
	resp, err := suite.srv.GetOriginalURL(context.Background(), &pb.GetOriginalURLRequest{
		ShortUrl: "iXAnV1m",
	})
	suite.Nil(err, "error should be nil")
	suite.Equal("https://google.com/helloworld", resp.OriginalUrl, "original url should be equal")
}

// In the hopes of trying to recreate somewhat double spending problem
// here the brave 100 requests arer made as goroutines.
func (suite *HandlersTestSuite) Test_getOriginalURL_MultipleRequests() {
	wg := &sync.WaitGroup{}
	numberOfRequests := 100
	wg.Add(numberOfRequests)
	for i := 0; i < numberOfRequests; i++ {
		go func() {
			defer wg.Done()
			_, err := suite.srv.GetOriginalURL(context.Background(), &pb.GetOriginalURLRequest{
				ShortUrl: "multipl",
			})
			suite.Nil(err, "error should be nil")
		}()
	}
	wg.Wait()

	urlResp, err := suite.db.GetURL(context.Background(), "multipl")
	suite.Nil(err, "error should be nil")

	suite.EqualValues(numberOfRequests, urlResp.Clicks, "clicks should be equal to number of requests")
}

func (suite *HandlersTestSuite) Test_getOriginalURL_NotFound() {
	resp, err := suite.srv.GetOriginalURL(context.Background(), &pb.GetOriginalURLRequest{
		ShortUrl: "aaaaaaa",
	})
	suite.Nil(resp, "response should be nil")
	suite.Require().NotNil(err, "error should not be nil")

	if errStatus, ok := status.FromError(err); ok {
		suite.Equal(errStatus.Code(), codes.NotFound, "error code should be not found (5)")
	} else {
		suite.Fail("error should be a GRPC status error")
	}
}

func (suite *HandlersTestSuite) Test_getURLStats_OK() {
	resp, err := suite.srv.GetURLStats(context.Background(), &pb.GetURLStatsRequest{
		ShortUrl: "eeeeeee",
	})
	suite.Nil(err, "error should be nil")
	suite.EqualValues(1, resp.Clicks, "clicks should be 1")
}
func (suite *HandlersTestSuite) Test_getURLStats_NotFound() {
	resp, err := suite.srv.GetURLStats(context.Background(), &pb.GetURLStatsRequest{
		ShortUrl: "aaaaaaa",
	})
	suite.Nil(resp, "response should be nil")
	suite.Require().NotNil(err, "error should not be nil")

	if errStatus, ok := status.FromError(err); ok {
		suite.Equal(errStatus.Code(), codes.NotFound, "error code should be not found (5)")
	} else {
		suite.Fail("error should be a GRPC status error")
	}
}

func (suite *HandlersTestSuite) Test_deleteURL_OK() {
	_, err := suite.srv.DeleteURL(context.Background(), &pb.DeleteURLRequest{
		ShortUrl: "nFziOn7",
	})
	suite.Nil(err, "error should be nil")

	// check database
	urlResp, err := suite.db.GetURL(context.Background(), "nFziOn7")
	suite.Nil(err, "error should be nil")
	suite.Nil(urlResp, "url should be nil")
}

func (suite *HandlersTestSuite) Test_deleteURL_NotFound() {
	resp, err := suite.srv.DeleteURL(context.Background(), &pb.DeleteURLRequest{
		ShortUrl: "aaaaaaa",
	})
	suite.Nil(resp, "response should be nil")
	suite.Require().NotNil(err, "error should not be nil")

	if errStatus, ok := status.FromError(err); ok {
		suite.Equal(errStatus.Code(), codes.NotFound, "error code should be not found (5)")
	} else {
		suite.Fail("error should be a GRPC status error")
	}
}

func (suite *HandlersTestSuite) prepareDatabase() {
	err := suite.db.AddURL(context.Background(), schema.URL{
		OriginalURL: "https://google.com/helloworld",
		ShortURLKey: "iXAnV1m",
	})
	suite.Nil(err, "error should be nil when adding a url")

	err = suite.db.AddURL(context.Background(), schema.URL{
		OriginalURL: "https://google.com/helloworld",
		ShortURLKey: "nFziOn7",
	})
	suite.Nil(err, "error should be nil when adding a url")

	err = suite.db.AddURL(context.Background(), schema.URL{
		OriginalURL: "https://google.com/helloworld",
		ShortURLKey: "eeeeeee",
	})
	suite.Nil(err, "error should be nil when adding a url")

	err = suite.db.IncrementClicks(context.Background(), "eeeeeee")
	suite.Nil(err, "error should be nil when incrementing clicks")

	// url for testing multiple requests
	err = suite.db.AddURL(context.Background(), schema.URL{
		OriginalURL: "https://google.com/helloworld",
		ShortURLKey: "multipl",
	})
	suite.Nil(err, "error should be nil when adding a url")
}

func (suite *HandlersTestSuite) TearDownSuite() {
	err := suite.db.FlushAll()
	suite.Nil(err, "error should be nil when flushing all")
}

func TestHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(HandlersTestSuite))
}
