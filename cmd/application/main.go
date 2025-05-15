package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BoanergesJunior/tracknme-challenge/cmd/application/setup"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/repository"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/usecase"
	handler "github.com/BoanergesJunior/tracknme-challenge/internal/http/handler"
	"github.com/joho/godotenv"
)

func main() {
	environment := os.Getenv("ENVIRONMENT")

	if environment != "production" {
		err := godotenv.Load()
		if err != nil {
			errMsg := fmt.Errorf("error loading .env file: %v", err)
			log.Fatalf("%v", errMsg)
			return
		}
	}

	database, err := setup.SetupDatabase(setup.Migrations{
		Path:   "migrations/up",
		Schema: "tracknme",
	})
	if err != nil {
		log.Fatal(err)
	}

	redis, err := setup.NewRedisClient()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	repository := repository.NewRepository(database, redis)

	uc := usecase.New(repository)

	h := handler.NewHandler(uc)

	if err := h.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
		log.Fatal(err)
	}
}
