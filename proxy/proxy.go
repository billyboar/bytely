package proxy

import (
	"bytes"
	"embed"
	"fmt"
	"net/http"
	"text/template"

	"github.com/billyboar/bytely/pb"
	"github.com/billyboar/bytely/proxy/config"
	"github.com/billyboar/bytely/proxy/docs"
	"github.com/billyboar/bytely/schema"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Webserver is a HTTP proxy for the GRPC service and also
// serves frontend HTML page.
type WebServer struct {
	grpcClient pb.BytelyServiceClient
	cfg        *config.Config
	engine     *gin.Engine
}

func NewWebserver(grpcClient pb.BytelyServiceClient, cfg *config.Config) *WebServer {
	return &WebServer{
		grpcClient: grpcClient,
		engine:     gin.Default(),
		cfg:        cfg,
	}
}

func (s *WebServer) ConfigureRoutes() {
	docs.SwaggerInfo.BasePath = "/api"
	s.engine.GET("/", s.indexHandler)
	s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiGroup := s.engine.Group("/api")
	{
		urlsGroup := apiGroup.Group("/urls")
		{
			urlsGroup.POST("", s.shortenURL)
			urlsGroup.GET("/:short_url", s.getURLStats)
			urlsGroup.DELETE("/:short_url", s.deleteURL)
		}
	}

	s.engine.GET("/:short_url", s.redirectURL)
}

//go:embed templates/index.html
var t embed.FS
var indexTemplate *template.Template

func (s *WebServer) LoadTemplate() error {
	var err error
	indexTemplate, err = template.ParseFS(t, "templates/index.html")
	return err
}

func (s *WebServer) indexHandler(ctx *gin.Context) {
	buf := &bytes.Buffer{}
	err := indexTemplate.Execute(buf, nil)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Internal server error")
		return
	}

	buf.WriteTo(ctx.Writer)
}

type shortenURLRequest struct {
	URL string `json:"url"`
}

type shortenURLResponse struct {
	ShortURL string `json:"short_url"`
}

// shortURL godoc
// @Summary receives a URL and returns a shortened URL.
// @Description This endpoint receives a URL and returns a shortened URL.
// @Accept json
// @Produce json
// @Param url body shortenURLRequest true "URL to shorten"
// @Success 200 {object} shortenURLResponse
// @Failure 400 {object} schema.ErrorResponse
// @Failure 500 {object} schema.ErrorResponse
// @Router /urls [post]
func (s *WebServer) shortenURL(c *gin.Context) {
	var req shortenURLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, schema.ErrorResponse{
			Code:    schema.ErrBadPayload,
			Message: "failed to parse request payload",
		})
		return
	}

	grpcReq := &pb.AddURLRequest{
		OriginalUrl: req.URL,
	}
	resp, err := s.grpcClient.AddURL(c.Request.Context(), grpcReq)
	if errStatus, ok := status.FromError(err); err != nil && ok {
		if errStatus.Code() == codes.InvalidArgument {
			c.JSON(http.StatusBadRequest, schema.ErrorResponse{
				Code:    schema.ErrBadPayload,
				Message: "Invalid URL",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, schema.ErrorResponse{
			Code:    schema.ErrGRPCServiceError,
			Message: "failed to shorten URL",
		})
		return
	}

	c.JSON(http.StatusOK, shortenURLResponse{
		ShortURL: fmt.Sprintf("http://%s/%s", s.cfg.Domain, resp.ShortUrl),
	})
}

type getURLStatsResponse struct {
	Clicks int `json:"clicks"`
}

// getURLStats godoc
// @Summary returns click count for a shortened URL.
// @Description This endpoint returns click count for a shortened URL.
// @Produce json
// @Param short_url path string true "Short URL to get stats"
// @Success 200 {object} getURLStatsResponse
// @Failure 404 {object} schema.ErrorResponse
// @Failure 500 {object} schema.ErrorResponse
// @Router /urls/{short_url} [get]
func (s *WebServer) getURLStats(c *gin.Context) {
	shortURL := c.Param("short_url")
	grpcReq := &pb.GetURLStatsRequest{
		ShortUrl: shortURL,
	}
	resp, err := s.grpcClient.GetURLStats(c.Request.Context(), grpcReq)
	if errStatus, ok := status.FromError(err); err != nil && ok {
		if errStatus.Code() == codes.NotFound {
			c.JSON(http.StatusNotFound, schema.ErrorResponse{
				Code:    schema.ErrNotFound,
				Message: "shortened URL not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, schema.ErrorResponse{
			Code:    schema.ErrGRPCServiceError,
			Message: "cannot fetch URL stats",
		})
	}

	c.JSON(http.StatusOK, getURLStatsResponse{
		Clicks: int(resp.Clicks),
	})
}

// deleteURL godoc
// @Summary deletes a shortened URL.
// @Description This endpoint deletes a shortened URL.
// @Produce json
// @Param short_url path string true "Short URL to delete"
// @Success 200
// @Failure 400 {object} schema.ErrorResponse
// @Failure 500 {object} schema.ErrorResponse
// @Router /urls/{short_url} [delete]
func (s *WebServer) deleteURL(c *gin.Context) {
	shortURL := c.Param("short_url")
	grpcReq := &pb.DeleteURLRequest{
		ShortUrl: shortURL,
	}
	_, err := s.grpcClient.DeleteURL(c.Request.Context(), grpcReq)
	if errStatus, ok := status.FromError(err); err != nil && ok {
		if errStatus.Code() == codes.NotFound {
			c.JSON(http.StatusNotFound, schema.ErrorResponse{
				Code:    schema.ErrNotFound,
				Message: "shortened URL not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, schema.ErrorResponse{
			Code:    schema.ErrGRPCServiceError,
			Message: "failed to delete URL",
		})
	}

	c.Status(http.StatusOK)
}

// redirectURL godoc
// @Summary find a correspoding URL for a short URL and redirects to it.
// @Description This endpoint find a correspoding URL for a short URL and redirects to it.
// @Param short_url path string true "Short URL to redirect to the original URL"
// @Success 302
// @Router /{short_url} [get]
func (s *WebServer) redirectURL(c *gin.Context) {
	shortURL := c.Param("short_url")
	grpcReq := &pb.GetOriginalURLRequest{
		ShortUrl: shortURL,
	}
	resp, err := s.grpcClient.GetOriginalURL(c.Request.Context(), grpcReq)
	if errStatus, ok := status.FromError(err); err != nil && ok {
		if errStatus.Code() == codes.NotFound {
			c.AbortWithError(http.StatusNotFound, errors.New("short_url is not found"))
			return
		}
		c.AbortWithError(http.StatusInternalServerError, errors.New("cannot redirect to URL"))
	}

	c.Redirect(http.StatusFound, resp.OriginalUrl)
}

func (s *WebServer) Start() error {
	return s.engine.Run(fmt.Sprintf(":%d", s.cfg.Port))
}
