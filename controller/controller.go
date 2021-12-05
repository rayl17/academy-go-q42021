package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/rayl17/academy-go-q42021/usecase"

	"github.com/gorilla/mux"
)

type ControllerInterface interface {
	GetPokemonByName(w http.ResponseWriter, r *http.Request)
	GetPokemonByID(w http.ResponseWriter, r *http.Request)
	PostPokemon(w http.ResponseWriter, r *http.Request)
	GetPokemonsConcurrently(w http.ResponseWriter, r *http.Request)
}

type PokemonController struct {
	usecase usecase.Usecaseinterface
}

func NewController(uc usecase.Usecaseinterface) *PokemonController {

	return &PokemonController{usecase: uc}

}

func (c *PokemonController) GetPokemonByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	defer r.Body.Close()

	if reflect.TypeOf(name).Kind() != reflect.String {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		encoderErr := json.NewEncoder(w).Encode(map[string]string{"error": "Name should be an string"})
		if encoderErr != nil {
			log.Println(encoderErr)
		}
	}

	pokemon, err := c.usecase.GetPokemonByName(name)

	if err != nil {

		w.WriteHeader(http.StatusNotFound)
		encoderErr := json.NewEncoder(w).Encode(map[string]string{"message": "Pokemon not found"})

		if encoderErr != nil {
			log.Println(encoderErr)
		}
	} else {
		w.Header().Add("Content-Type", "application/json")
		encodeErr := json.NewEncoder(w).Encode(*pokemon)
		if encodeErr != nil {
			log.Println(encodeErr)
		}
	}

}

func (c *PokemonController) GetPokemonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	defer r.Body.Close()

	///// validarr el id

	pokemon, err := c.usecase.GetPokemonByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		encoderErr := json.NewEncoder(w).Encode(map[string]string{"message": "Pokemon id  not found , Valid id`s 1-801"})
		if encoderErr != nil {
			log.Println(encoderErr)
		}
	} else {
		encoderErr := json.NewEncoder(w).Encode(*pokemon)
		if encoderErr != nil {
			log.Println(encoderErr)
		}
	}
}

func (c *PokemonController) PostPokemon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	if reflect.TypeOf(name).Kind() != reflect.String {
		w.WriteHeader(http.StatusBadRequest)
		encoderErr := json.NewEncoder(w).Encode(map[string]string{"error": "Name should be an string"})
		if encoderErr != nil {
			log.Println(encoderErr)
		}
	}

	message, err := c.usecase.SavePokemonByNameAPI(name)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		encoderErr := json.NewEncoder(w).Encode(map[string]string{"message": "Pokemon not found"})
		if encoderErr != nil {
			log.Println(encoderErr)
		}
	} else {
		encoderErr := json.NewEncoder(w).Encode(map[string]string{"message": message})
		if encoderErr != nil {
			log.Println(encoderErr)
		}
	}
}

func (c *PokemonController) GetPokemonsConcurrently(w http.ResponseWriter, r *http.Request) {
	types := r.URL.Query().Get("type")
	items := r.URL.Query().Get("items")
	itemsPW := r.URL.Query().Get("items_per_workers")
	defer r.Body.Close()

	itemsInt, errItems := strconv.Atoi(items)
	items_per_worker, errWorker := strconv.Atoi(itemsPW)

	if errItems != nil || errWorker != nil || items_per_worker == 0 || itemsInt == 0 || itemsInt%items_per_worker != 0 || types != "even" && types != "odd" || itemsInt < items_per_worker {

		err := ""

		if types != "even" && types != "odd" {
			err = "Type should be even or odd"

		} else if errWorker != nil || items_per_worker == 0 {
			err = "Items_per_workers should be a valid number more than 0"

		} else if errItems != nil || items_per_worker == 0 {

			err = "Items should be a valid number more than 0"

		} else if itemsInt%items_per_worker != 0 {
			err = "items_per_workers should be a multiple of items"
		} else if itemsInt < items_per_worker {
			err = "items should be greather than items_per_workers"
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		encoderErr := json.NewEncoder(w).Encode(map[string]string{"error": err})
		if encoderErr != nil {
			log.Println(encoderErr)
		}

	} else {
		pokemons, err := c.usecase.GetPokemonsConcurrently(types, itemsInt, items_per_worker)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			encoderErr := json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})

			if encoderErr != nil {
				log.Println(encoderErr)
			}
		} else {
			w.Header().Add("Content-Type", "application/json")
			encodeErr := json.NewEncoder(w).Encode(*pokemons)
			if encodeErr != nil {
				log.Println(encodeErr)
			}
		}
	}

}
