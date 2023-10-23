package app

import (
	"context"
	"fmt"

	"github.com/ScMofeoluwa/GatherGo/config"
	"github.com/ScMofeoluwa/GatherGo/internal/domain/user/repository"
	"github.com/ScMofeoluwa/GatherGo/internal/domain/user/service"
	"github.com/ScMofeoluwa/GatherGo/internal/infrastructure/database/sqlc"
	httpTransport "github.com/ScMofeoluwa/GatherGo/internal/transport/http"
	"github.com/ScMofeoluwa/GatherGo/internal/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Run(cfg config.Config) error {
	connPool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		return err
	}

	dbHandler := sqlc.NewSQLDbHandler(connPool)
	db := dbHandler.(*sqlc.SQLDbHandler).Queries

	userRepo := repository.NewUserRepository(db)
	jwtMaker := util.NewJWTMaker(cfg.AccessSecret, cfg.RefreshSecret)
	userService := service.NewUserService(userRepo, jwtMaker)

	handler := httpTransport.NewHandler(userService)

	// Start the HTTP server
	if err := handler.Serve(); err != nil {
		return err
	}

	fmt.Println("Server started successfully")
	return nil
}
