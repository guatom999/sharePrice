package sharePriceUseCases

import (
	sharePb "github.com/guatom999/sharePrice/sharePrice/sharePricePb"
)

type SharePriceUseCase interface {
	Test() (*sharePb.Test, error)
}
