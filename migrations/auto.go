package main

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"restapi-shop/configs"
	"restapi-shop/internal/user"
	"restapi-shop/internal/item"
	"restapi-shop/internal/review"
	"restapi-shop/internal/role"
	"restapi-shop/internal/cart"
)

func FormatDSN(db_conf *configs.DatabaseConnConfig) string {
        return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
                db_conf.Address,
                db_conf.Username,
                db_conf.Password,
                db_conf.DBName,
                db_conf.Port,
        )
}

func main() {
	cfg, err := configs.Load(os.Getenv("CONFIG_PATH"))
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(FormatDSN(&cfg.DatabaseConnConfig)), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&user.User{},
		&item.Item{},
		&review.Review{},
		&role.Role{},
		&cart.Cart{},
	)
}
