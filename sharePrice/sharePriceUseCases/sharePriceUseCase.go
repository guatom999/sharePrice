package sharePriceUseCases

import (
	"context"
	"log"
	"strconv"

	"github.com/guatom999/sharePrice/pkg/scrapper"
	"github.com/guatom999/sharePrice/sharePrice/sharePriceEntity"
	sharePb "github.com/guatom999/sharePrice/sharePrice/sharePricePb"
	"github.com/guatom999/sharePrice/sharePrice/sharePriceRepositories"
)

type sharePriceUseCase struct {
	sharePriceRepo sharePriceRepositories.SharePriceRepository
}

func NewSharePriceUseCase(sharePriceRepo sharePriceRepositories.SharePriceRepository) SharePriceUseCase {
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

func (u *sharePriceUseCase) AddTrackShare(pctx context.Context, req *sharePriceEntity.ShareTrackReq) (*sharePriceEntity.SharePrice, error) {

	result, err := u.SharePriceSearch(pctx, req.Name)

	if err != nil {
		log.Println("error: error is :", err)
		return nil, err
	}

	if err := u.sharePriceRepo.InsertSharePrice(pctx, result.Name, result.Price); err != nil {
		log.Println("error: error is :", err)
		return nil, err
	}

	return &sharePriceEntity.SharePrice{
		Name:  result.Name,
		Price: result.Price,
	}, nil
}
