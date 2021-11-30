package router

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rayl17/academy-go-q42021/controller"
	parameters "github.com/rayl17/academy-go-q42021/global"
)

func NewRouter(c controller.ControllerInterface) *http.Server {

	router := mux.NewRouter()

	router.HandleFunc("/pokemon/{name}", c.GetPokemonByName).Methods("GET")
	router.HandleFunc("/pokemonid/{id}", c.GetPokemonById).Methods("GET")
	router.HandleFunc("/pokemon/{name}", c.PostPokemon).Methods("POST")

	srv := http.Server{
		Handler:      router,
		Addr:         parameters.ServerConfig.Addres,
		WriteTimeout: time.Duration(parameters.ServerConfig.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(parameters.ServerConfig.WriteTimeout) * time.Second,
	}

	return &srv
}
