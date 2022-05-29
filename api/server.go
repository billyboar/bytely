package api

import (
	"fmt"
	"net"

	"github.com/billyboar/bytely/api/config"
	"github.com/billyboar/bytely/internal/storage"
	"github.com/billyboar/bytely/pb"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedBytelyServiceServer
	db  storage.Storage
	cfg *config.Config
}

func NewServer(db storage.Storage, cfg *config.Config) *Server {
	return &Server{
		db:  db,
		cfg: cfg,
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.Port))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBytelyServiceServer(grpcServer, s)
	return grpcServer.Serve(listener)
}

func (s *Server) Shutdown() {
	s.db.Close()
}
