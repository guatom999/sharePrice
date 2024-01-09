package sharePriceRepositories

import "context"

type SharePriceRepository interface {
	IsShareOnTracker(pctx context.Context, shareSymbol string) bool
	InsertSharePrice(pctx context.Context, shareSymbol string, sharePrice float64) error
	UpdateSharePrice(pctx context.Context, shareSymbol string) error
}
