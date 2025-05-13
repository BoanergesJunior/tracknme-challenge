package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BoanergesJunior/tracknme-challenge/cmd/application/setup"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/usecase"
	handler "github.com/BoanergesJunior/tracknme-challenge/internal/http/handler"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		errMsg := fmt.Errorf("error loading .env file: %v", err)
		log.Fatalf("%v", errMsg)
		return
	}

	database, err := setup.SetupDatabase(setup.Migrations{
		Path:   "migrations/up",
		Schema: "tracknme",
	})
	if err != nil {
		log.Fatal(err)
	}

	uc := usecase.New(database)

	h := handler.NewHandler(&uc)

	if err := h.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
		log.Fatal(err)
	}
}
