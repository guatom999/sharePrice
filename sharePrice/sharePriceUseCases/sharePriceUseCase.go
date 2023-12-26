package sharePriceUseCases

import (
	"context"
	"log"
	"strconv"

	"github.com/guatom999/sharePrice/pkg/scrapper"
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

func (u *sharePriceUseCase) SharePriceSearch(ctx context.Context, shareSymbol string) (*sharePb.SharePriceRes, error) {

	result, err := scrapper.Scrapper(shareSymbol)
	if err != nil {
		log.Println("error: error is :", err)
		return nil, err
	}

	resultFloat, err := strconv.ParseFloat(result, 32)
	if err != nil {
		return nil, err
	}

	return &sharePb.SharePriceRes{
		Name:  shareSymbol,
		Price: resultFloat,
	}, nil
}
