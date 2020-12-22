package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lcslucas/micro-service/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo .ENV: %v", err)
	}

	url_app := fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
	api.Run(url_app)
}
