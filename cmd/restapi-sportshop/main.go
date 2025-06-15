package main

import (
	"fmt"
	"os"
	"net/http"
	"log/slog"

	"restapi-sportshop/configs"
	"restapi-sportshop/pkg/db"
	"restapi-sportshop/pkg/slogpretty"
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

	logger := setupPrettySlog()

	logger.Debug(spew.Sdump(cfg))

	db, err := db.NewDb(cfg)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	itemRepo := item.NewItemRepository(db)
	userRepo := user.NewUserRepository(db)

	authService := auth.NewAuthService(userRepo)

	smux := http.NewServeMux()

	_ = auth.NewAuthHandler(smux, auth.AuthHandlerDeps{
		AuthService: authService,
		Logger: logger,
	})
	_ = user.NewUserHandler(smux, user.UserHandlerDeps{
		UserRepository: userRepo,
	})
	_ = item.NewItemHandler(smux, item.ItemHandlerDeps{
		ItemRepository: itemRepo,
		Logger: logger,
	})

	srv := http.Server{
		Addr: fmt.Sprintf(":%d", cfg.HTTPServerConfig.Port),
		Handler: smux,
	}

	logger.Info("Starting the server",
		"port", cfg.HTTPServerConfig.Port,
	)

	err = srv.ListenAndServe()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
