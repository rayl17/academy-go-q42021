package service_csv

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
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
}

type PokemonService struct {
	servicePath string
}

func NewPokemonService(path string) (*PokemonService, error) {

	return &PokemonService{
		servicePath: path,
	}, nil
}

func (ps *PokemonService) GetPokemonByName(name string) (*model.Pokemon, error) {

	for _, pokemon := range ps.getPokemons() {
		if pokemon.Name == name {
			return &pokemon, nil
		}
	}

	return nil, errors.New("Pokemon not found ")
}

func (ps *PokemonService) GetPokemonByID(id string) (*model.Pokemon, error) {

	for _, pokemon := range ps.getPokemons() {
		if pokemon.Id == id {
			return &pokemon, nil
		}
	}

	return nil, errors.New("Pokemon not found ")
}

func (p *PokemonService) getPokemons() []model.Pokemon {
	fmt.Print(p.servicePath)
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

func (ps *PokemonService) SearchPokemonApi(name string) (string, error) {
	res, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + name)

	if err != nil {
		return "", err

	}

	res2, err2 := savePokemonCSV(res)

	if err2 != nil {
		return "", err2
	}

	return res2, nil

}

func savePokemonCSV(res *http.Response) (string, error) {

	defer res.Body.Close()

	content, error1 := ioutil.ReadAll(res.Body)

	if error1 != nil {
		return "", error1
	}

	var pokemon = new(model.Poke)

	err2 := json.Unmarshal(content, &pokemon)

	if err2 != nil {
		return "", err2
	}

	file, err3 := os.OpenFile(parameters.CsvPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err3 != nil {
		return "", err3
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
