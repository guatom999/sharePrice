package database

import (
	"fmt"

	"github.com/guatom999/sharePrice/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDB struct {
	Db *gorm.DB
}

func NewPostgresDB(cfg *config.Config) Database {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.Db.Host,
		cfg.Db.User,
		cfg.Db.Password,
		cfg.Db.DbName,
		cfg.Db.Port,
		cfg.Db.SslMode,
		cfg.Db.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &postgresDB{Db: db}
}

func (p *postgresDB) GetDb() *gorm.DB {
	return p.Db
}
