package service_csv

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/raymundo/academy-go-q42021/model"
)

func GetPokemons() []model.Pokemon {
	f, err := os.Open("assets/pokemons.csv")
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

func formatPokemonList(data [][]string) []model.Pokemon {
	var pokemons []model.Pokemon

	for i, pokemon := range data {
		var dataPokemon []string = pokemon
		dataPokemon[1] = strings.ToLower(dataPokemon[1])

		if i > 0 {
			var Poke = model.Pokemon{
				Id:         dataPokemon[0],
				Name:       dataPokemon[1],
				Type1:      dataPokemon[2],
				Type2:      dataPokemon[3],
				Total:      dataPokemon[4],
				HP:         dataPokemon[5],
				Attack:     dataPokemon[6],
				Defense:    dataPokemon[7],
				SpAttack:   dataPokemon[8],
				SpDefense:  dataPokemon[9],
				Speed:      dataPokemon[10],
				Generation: dataPokemon[11],
				Legendary:  dataPokemon[12],
			}
			pokemons = append(pokemons, Poke)
		}

	}

	return pokemons
}
