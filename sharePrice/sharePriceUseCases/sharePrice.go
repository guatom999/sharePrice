package sharePriceUseCases

import (
	"context"

	sharePb "github.com/guatom999/sharePrice/sharePrice/sharePricePb"
)

type SharePriceUseCase interface {
	Test() (*sharePb.Test, error)
	SharePriceSearch(ctx context.Context, shareSymbol string) (*sharePb.SharePriceRes, error)
}
