package main

import (
	"log"

	"github.com/rayl17/academy-go-q42021/controller"
	parameters "github.com/rayl17/academy-go-q42021/global"
	"github.com/rayl17/academy-go-q42021/router"
	service_csv "github.com/rayl17/academy-go-q42021/service/csv"
	"github.com/rayl17/academy-go-q42021/usecase"
)

func main() {
	log.Print("Server is starting...")

	var err error

	pokemonService, err := service_csv.NewPokemonService(parameters.CsvPath)
	pokemonUsecase := usecase.NewUseCase(pokemonService)
	pokemonController := controller.NewController(pokemonUsecase)

	srv := router.NewRouter(pokemonController)
	log.Print("Server is up at port", parameters.ServerConfig.Addres)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
