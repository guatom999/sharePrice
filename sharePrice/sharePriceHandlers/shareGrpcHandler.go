package sharePriceHandlers

import (
	"context"

	sharePb "github.com/guatom999/sharePrice/sharePrice/sharePricePb"
	"github.com/guatom999/sharePrice/sharePrice/sharePriceUseCases"
)

type sharePriceGrpcHandler struct {
	sharePriceUsecase sharePriceUseCases.SharePriceUseCase
	sharePb.UnimplementedSharePriceGrpcServiceServer
}

func NewSharePriceGrpcHandler(sharePriceUsecase sharePriceUseCases.SharePriceUseCase) *sharePriceGrpcHandler {
	return &sharePriceGrpcHandler{
		sharePriceUsecase: sharePriceUsecase,
	}
}

func (g *sharePriceGrpcHandler) JustTest(context.Context, *sharePb.Test) (*sharePb.Test, error) {
	return g.sharePriceUsecase.Test()
}

func (g *sharePriceGrpcHandler) SharePriceSearch(ctx context.Context, req *sharePb.SharePriceReq) (*sharePb.SharePriceRes, error) {
	return g.sharePriceUsecase.SharePriceSearch(ctx, req.ShareSymbol)
}
