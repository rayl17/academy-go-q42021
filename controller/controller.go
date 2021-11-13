package controller

import (
	"bytes"
	"encoding/json"
	"net/http"

	service "github.com/raymundo/academy-go-q42021/service/csv"

	"github.com/gorilla/mux"
)

func GetPokemonByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	for _, pokemon := range service.GetPokemons() {
		if pokemon.Name == name {
			json.NewEncoder(w).Encode(pokemon)
			return
		}
	}

	var buffer bytes.Buffer
	buffer.WriteString(`{message:"Pokemon not found"}`)
	json.NewEncoder(w).Encode(map[string]string{"message": "Pokemon not found"})

}

func GetPokemonById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, pokemon := range service.GetPokemons() {
		if id == pokemon.Id {
			json.NewEncoder(w).Encode(pokemon)
			return
		}
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Pokemon id  not found , Valid id`s 1-801"})

}
