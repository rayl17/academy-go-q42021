package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/rayl17/academy-go-q42021/usecase"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
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

/// in this method recives the parameters from the request , get the vaiable name from the  request , verify is an string
/// then  search for given pokemon name and send the response back in case the pokemon exist
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

/// in this method recives the parameters from the request , get the vaiable ID from the  request , verify is an string
/// then  search for given pokemon name and send the response back in case the pokemon exist
func (c *PokemonController) GetPokemonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	defer r.Body.Close()

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

/// in this method recives the parameters from the request , get the vaiable name from the  request , verify is an string
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

/// in this method recives the parameters from the request , get the vquery parameters type , items , items_per_worker  validate the corect values
/// then it rerurs an error if it occurs or return a list op pokemons
func (c *PokemonController) GetPokemonsConcurrently(w http.ResponseWriter, r *http.Request) {
	types := r.URL.Query().Get("type")
	items := r.URL.Query().Get("items")
	itemsPW := r.URL.Query().Get("items_per_workers")
	defer r.Body.Close()
	w.Header().Add("Content-Type", "application/json")

	itemsInt, err := strconv.Atoi(items)
	if err != nil && itemsInt == 0 {
		err = errors.Wrap(err, "items should be an integer")
		w.WriteHeader(http.StatusBadRequest)
		encoderErr := json.NewEncoder(w).Encode(map[string]string{"error": "items should be an integer"})
		if err != nil {
			err = errors.Wrap(err, encoderErr.Error())
		}
		log.Println(err)

		return
	}

	items_per_worker, errWorker := strconv.Atoi(itemsPW)
	if errWorker != nil && items_per_worker == 0 {
		errWorker = errors.Wrap(errWorker, "items_per_workers should be an integer grather than 0")
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(map[string]string{"error": "items_per_workers should be an integer grather than 0"})
		if err != nil {
			err = errors.Wrap(errWorker, err.Error())
		}
		log.Println(err)
		return
	}

	switch types {
	case "":
		err = errors.Wrap(err, "type should  not be empty")
	case "odd", "even":
		break
	default:
		err = errors.Wrap(err, "type should be odd or even")
	}

	if err != nil {
		err = errors.Wrap(err, "Error getting type param")
		w.WriteHeader(http.StatusBadRequest)
		encoderErr := json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		if encoderErr != nil {
			err = errors.Wrap(err, encoderErr.Error())
		}
		log.Println(err)
		return
	}

	pokemons, err := c.usecase.GetPokemonsConcurrently(types, itemsInt, items_per_worker)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		encoderErr := json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})

		if encoderErr != nil {
			err = errors.Wrap(err, encoderErr.Error())
		}
		log.Println(err)
		return
	}

	encodeErr := json.NewEncoder(w).Encode(*pokemons)
	if encodeErr != nil {
		log.Println(encodeErr)
	}

}
