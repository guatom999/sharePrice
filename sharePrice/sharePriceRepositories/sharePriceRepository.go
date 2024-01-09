package sharePriceRepositories

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/guatom999/sharePrice/pkg/scrapper"
	"github.com/guatom999/sharePrice/sharePrice/sharePriceEntity"
	"github.com/guatom999/sharePrice/utils"
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

	tx := r.db.Create(&data)
	if tx.Error != nil {
		log.Errorf("Error: Create Failed: %v", tx.Error)
		return tx.Error
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

func (r *sharePriceRepository) UpdateSharePrice(pctx context.Context, shareSymbol string) error {

	// oldshare := new(sharePriceEntity.SharePrice)
	oldshare := sharePriceEntity.SharePrice{}

	tx := r.db.Where("name = ?", strings.ToUpper(shareSymbol)).First(&oldshare)
	if tx.Error != nil {
		log.Errorf("Error: Search Failed: %v", tx.Error)
		return tx.Error
	}

	fmt.Println("old share is ", oldshare)

	price, err := scrapper.Scrapper(oldshare.Name)

	if err != nil {
		return err
	}

	priceFloat, err := strconv.ParseFloat(price, 32)

	if err != nil {
		return errors.New("error: convert string price to float failed")
	}

	newShare := sharePriceEntity.SharePrice{
		Id:        oldshare.Id,
		Name:      oldshare.Name,
		Price:     priceFloat,
		UpdatedAt: utils.LoadLocalTime(),
	}

	tx = r.db.Save(&newShare)
	if tx.Error != nil {
		log.Errorf("Error: Update SharePrice Failed: %v", tx.Error)
		return tx.Error
	}

	return nil

}
