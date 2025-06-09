package main

import (
	"fmt"
	"os"
	"net/http"

	"restapi-sportshop/configs"
	"restapi-sportshop/pkg/db"
	"restapi-sportshop/internal/user"
	"restapi-sportshop/internal/auth"
	"restapi-sportshop/internal/item"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	config_path := os.Getenv("CONFIG_PATH")
	cfg, err := configs.Load(config_path)

	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	spew.Dump(cfg)

	db, err := db.NewDb(cfg)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	itemRepo := item.NewItemRepository(db)

	smux := http.NewServeMux()

	_ = auth.NewAuthHandler(smux, auth.AuthHandlerDeps{})
	_ = user.NewUserHandler(smux, user.UserHandlerDeps{})
	_ = item.NewItemHandler(smux, item.ItemHandlerDeps{
		ItemRepository: itemRepo,
	})

	srv := http.Server{
		Addr: fmt.Sprintf(":%d", cfg.HTTPServerConfig.Port),
		Handler: smux,
	}

	fmt.Printf("Listening on the port %d\n", cfg.HTTPServerConfig.Port)
	err = srv.ListenAndServe()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
