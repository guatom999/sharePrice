package sharePriceRepositories

import (
	"context"

	"github.com/guatom999/sharePrice/sharePrice/sharePriceEntity"
	"gorm.io/gorm"
)

type sharePriceRepositoey struct {
	db *gorm.DB
}

func NewSharePriceRepository(db *gorm.DB) SharePriceRepositoey {
	return &sharePriceRepositoey{db: db}
}

func (r *sharePriceRepositoey) InsertSharePrice(pctx context.Context, shareSymbol string, sharePrice float64) error {

	data := &sharePriceEntity.SharePrice{
		Name:  shareSymbol,
		Price: sharePrice,
	}

	return nil
}
