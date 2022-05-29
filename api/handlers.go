package api

import (
	"context"
	"math/rand"

	"github.com/billyboar/bytely/pb"
	"github.com/billyboar/bytely/schema"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const defaultURLLength = 7

func (s *Server) generateShortURL(ctx context.Context) (string, error) {
	shortURL := make([]byte, defaultURLLength)
	for i := range shortURL {
		shortURL[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}

	u, err := s.db.GetURL(ctx, string(shortURL))
	if err != nil {
		return "", err
	}

	// found an url with same short_url in database already
	if u == nil && err == nil {
		return string(shortURL), nil
	}

	return s.generateShortURL(ctx)
}

func (s *Server) AddURL(ctx context.Context, req *pb.AddURLRequest) (*pb.AddURLResponse, error) {
	// generate short_url
	shortURL, err := s.generateShortURL(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to generate short url")
		return nil, status.Error(codes.Internal, "failed to generate short url")
	}

	// store url
	url := schema.URL{
		OriginalURL: req.OriginalUrl,
		ShortURLKey: shortURL,
	}

	if err := url.Validate(); err != nil {
		log.Debug().Err(err).Msg("validation failed")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := s.db.AddURL(ctx, url); err != nil {
		log.Error().Err(err).Msg("cannot store url")
		return nil, status.Error(codes.Internal, "failed to add url")
	}

	return &pb.AddURLResponse{ShortUrl: url.ShortURLKey}, nil
}

func (s *Server) GetOriginalURL(ctx context.Context, req *pb.GetOriginalURLRequest) (*pb.GetOriginalURLResponse, error) {
	url, err := s.db.GetURL(ctx, req.ShortUrl)
	if url == nil && err == nil {
		return nil, status.Error(codes.NotFound, "url not found")
	}
	if err != nil {
		log.Error().Err(err).Msg("failed to get url")
		return nil, status.Error(codes.Internal, "failed to fetch url")
	}

	// use background context because redirect will close off connection
	// before database can be updated
	if err := s.db.IncrementClicks(context.Background(), req.ShortUrl); err != nil {
		log.Error().Err(err).Msg("failed to increment visits")
	}

	return &pb.GetOriginalURLResponse{OriginalUrl: url.OriginalURL}, nil
}

func (s *Server) GetURLStats(ctx context.Context, req *pb.GetURLStatsRequest) (*pb.GetURLStatsResponse, error) {
	url, err := s.db.GetURL(ctx, req.ShortUrl)
	if url == nil && err == nil {
		return nil, status.Error(codes.NotFound, "url not found")
	}
	if err != nil {
		log.Error().Err(err).Msg("failed to get url")
		return nil, status.Error(codes.Internal, "failed to fetch url")
	}

	return &pb.GetURLStatsResponse{Clicks: int64(url.Clicks)}, nil
}

func (s *Server) DeleteURL(ctx context.Context, req *pb.DeleteURLRequest) (*emptypb.Empty, error) {
	ok, err := s.db.DeleteURL(ctx, req.ShortUrl)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to delete url")
	}
	if !ok {
		return nil, status.Error(codes.NotFound, "url not found")
	}

	return &emptypb.Empty{}, nil
}
