package service_csv

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/rayl17/academy-go-q42021/model"

	parameters "github.com/rayl17/academy-go-q42021/global"
)

type ServiceInterface interface {
	GetPokemonByName(name string) (*model.Pokemon, error)
	GetPokemonByID(id string) (*model.Pokemon, error)
	getPokemons() []model.Pokemon
	SearchPokemonApi(name string) (string, error)
	GetPokemonsByTypes(types string, items int, items_per_worker int) (*[]model.Pokemon, error)
}

type PokemonService struct {
	servicePath string
}

func NewPokemonService(path string) (*PokemonService, error) {

	return &PokemonService{
		servicePath: path,
	}, nil
}

// this method recives a name and return the pokemon with that name
func (ps *PokemonService) GetPokemonByName(name string) (*model.Pokemon, error) {

	for _, pokemon := range ps.getPokemons() {
		if pokemon.Name == name {
			return &pokemon, nil
		}
	}

	return nil, errors.New("Pokemon not found ")
}

/// This method recives a id to look for and returns the pokemon with that id ,
func (ps *PokemonService) GetPokemonByID(id string) (*model.Pokemon, error) {

	for _, pokemon := range ps.getPokemons() {
		if pokemon.Id == id {
			return &pokemon, nil
		}
	}

	return nil, errors.New("Pokemon not found")
}

// this fmethod reads thecvs file and   return a list of pokemons correctli formated  from the csv file
func (p *PokemonService) getPokemons() []model.Pokemon {
	f, err := os.Open(p.servicePath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	pokemons := formatPokemonList(data)
	return pokemons
}

/// this mehod  make a get request to the api , call the savePokemonCSV function and return the response
/// if is succesfully or not
func (ps *PokemonService) SearchPokemonApi(name string) (string, error) {
	res, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + name)

	if err != nil {
		return "", err

	}

	res2, err := savePokemonCSV(res)

	if err != nil {
		return "", err
	}
	return res2, nil
}

/// this function recives the response from the api and  and  add the pokemon to the csv file row
func savePokemonCSV(res *http.Response) (string, error) {

	defer res.Body.Close()

	content, error1 := ioutil.ReadAll(res.Body)

	if error1 != nil {
		return "", error1
	}

	var pokemon = new(model.Poke)

	err := json.Unmarshal(content, &pokemon)

	if err != nil {
		return "", err
	}

	file, err := os.OpenFile(parameters.CsvPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return "", err
	}

	defer file.Close()
	var type1 string = "no type"
	var type2 string = "no type"
	if len(pokemon.Types) == 1 {
		type1 = pokemon.Types[0].Type.Name
	} else if len(pokemon.Types) == 2 {
		type1 = pokemon.Types[0].Type.Name
		type2 = pokemon.Types[1].Type.Name
	}

	var data [][]string
	var row []string = []string{strconv.Itoa(pokemon.ID), pokemon.Name, type1, type2, "0", "0", "0", "0", "0", "0", "0", "0", "false"}
	data = append(data, row)
	q := csv.NewWriter(file)

	err4 := q.WriteAll(data)

	if err4 != nil {
		return "", err4
	}

	return "Saved succesfully", nil

}

/// this method   returns a list of pokemons by type and requred parameters , apliying the worker pool
func (p *PokemonService) GetPokemonsByTypes(types string, items int, items_per_worker int) (*[]model.Pokemon, error) {

	var quantity_workers = items / items_per_worker

	jobs := make(chan model.Pokemon, 1000)
	results := make(chan model.Pokemon, items)
	pokemons := make([]model.Pokemon, 0)
	pokemonList := p.getPokemons()

	for w := 1; w <= quantity_workers; w++ {
		go worker(w, jobs, results, types, items_per_worker)
	}

	for _, a := range pokemonList {
		jobs <- a
	}
	close(jobs)

	for a := 1; a <= items; a++ {
		poke, _ := <-results
		pokemons = append(pokemons, poke)
	}
	close(results)

	return &pokemons, nil
}

// this functin is a worker that recives a pokemon by jobs and return a pokemon by results  depending on tipo and items_per_worker
func worker(id int, jobs <-chan model.Pokemon, results chan<- model.Pokemon, tipo string, items_per_worker int) {

	var count = 0
	for j := range jobs {
		pokemonID, _ := strconv.Atoi(j.Id)
		if items_per_worker == count {
			break
		}
		if tipo == "even" && pokemonID%2 == 0 {
			count++
			results <- j

		} else if tipo == "odd" && pokemonID%2 != 0 {
			count++
			results <- j
		}
	}
}

/// This function recives a list  [][]string and returns a list of model.Pokemon correctly formatted
func formatPokemonList(data [][]string) []model.Pokemon {
	var pokemons []model.Pokemon
	var poke model.Pokemon

	for _, pokemon := range data[1:] {

		poke = model.Pokemon{
			Id:         pokemon[0],
			Name:       strings.ToLower(pokemon[1]),
			Type1:      pokemon[2],
			Type2:      pokemon[3],
			Total:      pokemon[4],
			HP:         pokemon[5],
			Attack:     pokemon[6],
			Defense:    pokemon[7],
			SpAttack:   pokemon[8],
			SpDefense:  pokemon[9],
			Speed:      pokemon[10],
			Generation: pokemon[11],
			Legendary:  pokemon[12],
		}
		pokemons = append(pokemons, poke)

	}
	return pokemons
}
