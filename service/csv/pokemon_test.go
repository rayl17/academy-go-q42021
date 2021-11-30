package service_csv

import (
	"testing"

	parameters "github.com/rayl17/academy-go-q42021/global"
	"github.com/rayl17/academy-go-q42021/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mocks struct {
	mock.Mock
}

var pokemonData = []model.Pokemon{

	{
		Id:         "1",
		Name:       "Bulbasaur",
		Type1:      "Grass",
		Type2:      "Poison",
		Total:      "318",
		HP:         "45",
		Attack:     "49",
		Defense:    "49",
		SpAttack:   "65",
		SpDefense:  "65",
		Speed:      "45",
		Generation: "1",
		Legendary:  "False",
	},
}

func (mr mocks) GetPokemonByName(name string) (*model.Pokemon, error) {
	arg := mr.Called()
	return arg.Get(0).(*model.Pokemon), arg.Error(1)
}
func TestGetPokemonByName(t *testing.T) {

	t.Run("bulbasaur", func(t *testing.T) {

		mock := mocks{}
		mock.On("GetPokemonByName").Return(pokemonData, nil)

		service, error := NewPokemonService(parameters.CsvPath)

		assert.Nil(t, error)

		data, err := service.GetPokemonByName("bulbasaur")

		assert.Nil(t, err)

		assert.Equal(t, pokemonData, data)

	})
}

func TestGetPokemonByName2(t *testing.T) {

	expected := "pikachu"

	service, err := NewPokemonService(parameters.CsvPath)

	pokemon, err1 := service.GetPokemonByName("pikachu")

	assert.Nil(t, err1)
	assert.Nil(t, err)

	assert.Equal(t, expected, pokemon.Name, "they should be equal")

}

func TestGetPokemonByID(t *testing.T) {

	expected := "bulbasaur"

	service, err := NewPokemonService(parameters.CsvPath)

	pokemon, err1 := service.GetPokemonByID("1")

	assert.Nil(t, err1)
	assert.Nil(t, err)

	assert.Equal(t, expected, pokemon.Name, "they should be equal")

}
