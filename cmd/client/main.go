package main

import (
	"github.com/billyboar/bytely/pb"
	"github.com/billyboar/bytely/proxy"
	"github.com/billyboar/bytely/proxy/config"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Panic().Err(err).Msg("failed to load config")
	}
	conn, err := grpc.Dial(cfg.GRPCServerAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to the grpc server")
	}

	grpcClient := pb.NewBytelyServiceClient(conn)

	srv := proxy.NewWebserver(grpcClient, cfg)
	srv.ConfigureRoutes()
	if err := srv.LoadTemplate(); err != nil {
		log.Fatal().Err(err).Msg("cannot load http template")
	}

	if err := srv.Start(); err != nil {
		log.Fatal().Err(err).Msg("cannot start the HTTP server")
	}
}
