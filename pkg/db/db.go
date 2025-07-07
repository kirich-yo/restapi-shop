package db

import (
	"fmt"

	"restapi-sportshop/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func FormatDSN(db_conf *configs.DatabaseConnConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		db_conf.Address,
		db_conf.Username,
		db_conf.Password,
		db_conf.DBName,
		db_conf.Port,
	)
}

func NewDb(conf *configs.Config) (*Db, error) {
	dsn := FormatDSN(&conf.DatabaseConnConfig)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		return nil, err
	}

	return &Db{DB: db}, nil
}
