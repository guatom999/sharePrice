package sharePriceRepositories

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/guatom999/sharePrice/sharePrice/sharePriceEntity"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type sharePriceRepository struct {
	db *gorm.DB
}

func NewSharePriceRepository(db *gorm.DB) SharePriceRepository {
	return &sharePriceRepository{db: db}
}

func (r *sharePriceRepository) InsertSharePrice(pctx context.Context, shareSymbol string, sharePrice float64) error {

	data := &sharePriceEntity.SharePrice{
		Name:  strings.ToUpper(shareSymbol),
		Price: sharePrice,
	}

	if !r.IsShareOnTracker(pctx, shareSymbol) {
		return errors.New("error: share is already on track")
	}

	result := r.db.Create(&data)
	if result.Error == nil {
		log.Errorf("Error: Create Failed: %v", result.Error)
		return result.Error
	}

	return nil
}

func (r *sharePriceRepository) IsShareOnTracker(pctx context.Context, shareSymbol string) bool {

	share := make([]sharePriceEntity.SharePrice, 0)

	r.db.Where("name = ?", shareSymbol).Find(&share)

	if share != nil {
		fmt.Println("Error: share is already on track")
		return true
	}

	return false
}
