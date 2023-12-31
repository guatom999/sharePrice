package sharePriceUseCases

import (
	"context"

	"github.com/guatom999/sharePrice/sharePrice/sharePriceEntity"
	sharePb "github.com/guatom999/sharePrice/sharePrice/sharePricePb"
)

type SharePriceUseCase interface {
	Test() (*sharePb.Test, error)
	SharePriceSearch(ctx context.Context, shareSymbol string) (*sharePb.SharePriceRes, error)
	AddTrackShare(pctx context.Context, req *sharePriceEntity.ShareTrackReq) (*sharePriceEntity.SharePrice, error)
}
