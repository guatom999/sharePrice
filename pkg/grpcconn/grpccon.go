package grpcconn

import (
	"errors"
	"log"
	"net"

	"github.com/guatom999/sharePrice/config"
	sharePricePb "github.com/guatom999/sharePrice/sharePrice/sharePricePb"
	"google.golang.org/grpc"
)

type (
	GrcpClientFactoryHandler interface {
		SharePrice() sharePricePb.SharePriceGrpcServiceClient
	}

	grpcClientFactory struct {
		client *grpc.ClientConn
	}
)

func (g *grpcClientFactory) SharePrice() sharePricePb.SharePriceGrpcServiceClient {
	return sharePricePb.NewSharePriceGrpcServiceClient(g.client)
}

func NewGrpcClient(host string) (GrcpClientFactoryHandler, error) {

	clientConn, err := grpc.Dial(host)

	if err != nil {
		log.Printf("Error: Grpc client connection failed:%v", err)
		return nil, errors.New("error: grpc client connection failed:%")
	}

	defer clientConn.Close()

	return &grpcClientFactory{client: clientConn}, nil
}

func NewGrpcServer(host string, cfg *config.Config) (*grpc.Server, net.Listener) {

	grpcServer := grpc.NewServer()

	lis, err := net.Listen("tcp", host)

	if err != nil {
		log.Printf("Error: Failed to listen: %v", err)
	}

	return grpcServer, lis
}
