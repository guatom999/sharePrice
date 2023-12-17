package sharePriceRepositories

import "gorm.io/gorm"

type sharePriceRepositoey struct {
	db *gorm.DB
}

func NewSharePriceRepository(db *gorm.DB) SharePriceRepositoey {
	return &sharePriceRepositoey{db: db}
}
