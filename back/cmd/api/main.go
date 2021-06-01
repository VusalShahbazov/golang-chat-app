package main

import (
	"back/internal/api"
	"github.com/joho/godotenv"
	"log"

)

func main()  {
	err := godotenv.Load()

	apiServer := api.NewApiServer()

	err = apiServer.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}