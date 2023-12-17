package server

import (
	"log"

	"github.com/guatom999/sharePrice/pkg/grpcconn"
	"github.com/guatom999/sharePrice/sharePrice/sharePriceHandlers"
	sharePb "github.com/guatom999/sharePrice/sharePrice/sharePricePb"
	"github.com/guatom999/sharePrice/sharePrice/sharePriceRepositories"
	"github.com/guatom999/sharePrice/sharePrice/sharePriceUseCases"
)

func (s *server) sharePriceServer() {
	sharePriceRepository := sharePriceRepositories.NewSharePriceRepository(s.db)
	sharePriceUseCase := sharePriceUseCases.NewSharePriceUseCase(sharePriceRepository)
	sharePriceHandler := sharePriceHandlers.NewSharePriceHandler(s.cfg, sharePriceUseCase)
	sharePriceGrpcHandler := sharePriceHandlers.NewSharePriceGrpcHandler(sharePriceUseCase)

	// _ = sharePriceHandler

	go func() {

		grpcServer, listen := grpcconn.NewGrpcServer("0.0.0.0:1425", s.cfg)

		sharePb.RegisterSharePriceGrpcServiceServer(grpcServer, sharePriceGrpcHandler)

		log.Println("player Grpc server listening on: 0.0.0.0:1425")

		grpcServer.Serve(listen)

	}()

	router := s.app.Group("/shareprice_v1")

	router.GET("/test", sharePriceHandler.Test)

}
