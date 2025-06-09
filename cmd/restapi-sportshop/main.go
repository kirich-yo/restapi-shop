package main

import (
	"fmt"
	"os"
	"net/http"

	"restapi-sportshop/configs"
	"restapi-sportshop/internal/user"
	"restapi-sportshop/internal/auth"

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

	smux := http.NewServeMux()

	_ = user.NewUserHandler(smux, user.UserHandlerDeps{})
	_ = auth.NewAuthHandler(smux, auth.AuthHandlerDeps{})

	srv := http.Server{
		Addr: fmt.Sprintf(":%d", cfg.HTTPServerConfig.Port),
		Handler: smux,
	}

	err = srv.ListenAndServe()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
