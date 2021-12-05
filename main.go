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
	println("Server is starting...")

	pokemonService, _ := service_csv.NewPokemonService(parameters.CsvPath)

	pokemonUsecase := usecase.NewUseCase(pokemonService)

	pokemonController := controller.NewController(pokemonUsecase)

	srv := router.NewRouter(pokemonController)
	println("Server is up at port", parameters.ServerConfig.Addres)

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
