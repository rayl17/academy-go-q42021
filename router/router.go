package router

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/raymundo/academy-go-q42021/controller"
	"github.com/raymundo/academy-go-q42021/global"
)

func NewRouter() *http.Server {
	router := mux.NewRouter()

	router.HandleFunc("/pokemon/{name}", controller.GetPokemonByName).Methods("GET")
	router.HandleFunc("/pokemonid/{id}", controller.GetPokemonById).Methods("GET")

	srv := http.Server{
		Handler:      router,
		Addr:         global.ServerConfig.Addres,
		WriteTimeout: time.Duration(global.ServerConfig.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(global.ServerConfig.WriteTimeout) * time.Second,
	}

	return &srv
}
