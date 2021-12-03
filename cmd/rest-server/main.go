package main

import (
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/KrisCatDog/go-standard-layered-boilerplate/internal/config"
	"github.com/KrisCatDog/go-standard-layered-boilerplate/internal/server"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

//go:embed static
var static embed.FS

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}

	pool, err := config.NewPostgreSQL()
	if err != nil {
		return err
	}

	srv, err := server.New(server.Config{
		Address: fmt.Sprintf(":%s", os.Getenv("APP_PORT")),
		DB:      pool,
		Logger:  logger,
		Static:  static,
	})
	if err != nil {
		return err
	}

	if err := srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
