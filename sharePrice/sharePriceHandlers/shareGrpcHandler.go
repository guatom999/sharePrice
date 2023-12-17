package sharePriceHandlers

import (
	"context"

	sharePb "github.com/guatom999/sharePrice/sharePrice/sharePricePb"
	sharePrice "github.com/guatom999/sharePrice/sharePrice/sharePricePb"
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

func (g *sharePriceGrpcHandler) JustTest(context.Context, *sharePrice.Test) (*sharePb.Test, error) {
	return g.sharePriceUsecase.Test()
}
