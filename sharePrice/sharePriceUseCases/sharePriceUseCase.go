package sharePriceUseCases

import (
	sharePb "github.com/guatom999/sharePrice/sharePrice/sharePricePb"
	"github.com/guatom999/sharePrice/sharePrice/sharePriceRepositories"
)

type sharePriceUseCase struct {
	sharePriceRepo sharePriceRepositories.SharePriceRepositoey
}

func NewSharePriceUseCase(sharePriceRepo sharePriceRepositories.SharePriceRepositoey) SharePriceUseCase {
	return &sharePriceUseCase{sharePriceRepo: sharePriceRepo}
}

func (u *sharePriceUseCase) Test() (*sharePb.Test, error) {
	return &sharePb.Test{
		Message: "Test",
	}, nil
}
