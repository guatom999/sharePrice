package jobs

import (
	"context"
	"fmt"
	"time"

	"github.com/guatom999/sharePrice/sharePrice/sharePriceEntity"
	"github.com/guatom999/sharePrice/sharePrice/sharePriceRepositories"
	"github.com/guatom999/sharePrice/utils"
	"gorm.io/gorm"
)

type (
	JobSharePriceService interface {
		UpdateSharePriceJob(pctx context.Context, db *gorm.DB)
	}

	jobSharePrice struct {
		sharePriceRepo sharePriceRepositories.SharePriceRepository
		// db *gorm.DB
	}
)

func NewJobsSharePrice(sharePriceRepo sharePriceRepositories.SharePriceRepository) JobSharePriceService {
	return &jobSharePrice{sharePriceRepo: sharePriceRepo}
}

func (r *jobSharePrice) UpdateSharePriceJob(pctx context.Context, db *gorm.DB) {

	shares := []sharePriceEntity.SharePrice{}

	date := utils.GetLocalDate()
	for date.String() != "Saturday" || date.String() != "Sunday" {
		time.Sleep(10 * time.Second)
		// fmt.Println("Update Share Price")

		// tx := db.Order("id").Find(&share)
		tx := db.Find(&shares)
		if tx.Error != nil {
			fmt.Println(tx.Error)
			return
		}

		for i, share := range shares {
			fmt.Println("Update Share Price number:", i)
			r.sharePriceRepo.UpdateSharePrice(pctx, share.Name)
		}

		// fmt.Println(shares[0].Name)

	}
}
