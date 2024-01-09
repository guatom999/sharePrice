package main

import (
	"github.com/guatom999/sharePrice/config"
	"github.com/guatom999/sharePrice/database"
	"github.com/guatom999/sharePrice/sharePrice/sharePriceEntity"
)

func main() {
	cfg := config.GetConfig()

	db := database.NewPostgresDB(&cfg)

	sharePriceMigrate(db)
}

func sharePriceMigrate(db database.Database) {
	// db.GetDb().Migrator().CreateTable(&sharePriceEntity.SharePrice{})
	db.GetDb().CreateInBatches([]sharePriceEntity.SharePrice{
		{Name: "COM7", Price: 22.7},
		{Name: "CPF", Price: 18.8},
		{Name: "CPALL", Price: 18.8},
	}, 10)
}
