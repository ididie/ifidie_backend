package app

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ididie/ifidie_backend/internal/database/pg"
	"github.com/ididie/ifidie_backend/internal/services/auth/login"

	"github.com/spf13/viper"
)

func Run() error {
	errViper := initializeViper()
	if errViper != nil {
		return fmt.Errorf("app error on initializeViper: %w", errViper)
	}

	errPostgres := initializePostgres()
	if errPostgres != nil {
		return fmt.Errorf("app error on initializePostgres: %w", errPostgres)
	}

	runApiServer()

	return nil
}

func initializeViper() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		return fmt.Errorf("Error while reading config file: %w", err)
	}

	return nil
}

func initializePostgres() error {
	pgHost := viper.GetString("postgres.host")
	pgPort := viper.GetInt("postgres.port")
	pgDbName := viper.GetString("postgres.db_name")
	pgUser := viper.GetString("postgres.user")
	pgPassword := viper.GetString("postgres.password")

	_, err := pg.ConnectPostgres(context.Background(), fmt.Sprintf("postgresql://%s:%d/%s?user=%s&password=%s", pgHost, pgPort,
		pgDbName,
		pgUser,
		pgPassword))

	if err != nil {
		return fmt.Errorf("error on initializing postgres: %w", err)
	}

	return nil
}

func runApiServer() {
	apiHost := viper.GetString("api.host")
	apiPort := viper.GetInt("api.port")

	r := gin.Default()

	r.GET("/auth/login", login.Login)

	r.Run(fmt.Sprintf("%s:%d", apiHost, apiPort))
}
