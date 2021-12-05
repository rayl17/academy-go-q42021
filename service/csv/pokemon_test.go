package service_csv

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/rayl17/academy-go-q42021/mocks"
	"github.com/rayl17/academy-go-q42021/model"
	"github.com/stretchr/testify/assert"
)

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

var prueba = []model.Pokemon{
	{
		Id:         "1",
		Name:       "algo",
		Type1:      "ewrg",
		Type2:      "reg",
		Total:      "erg",
		HP:         "erg",
		Attack:     "reg",
		Defense:    "rew",
		SpAttack:   "ret",
		SpDefense:  "trw",
		Speed:      "ret",
		Generation: "ret",
		Legendary:  "wer",
	},
	{
		Id:         "2",
		Name:       "picaku",
		Type1:      "asd",
		Type2:      "fdsg",
		Total:      "dfgs",
		HP:         "dsfg",
		Attack:     "dfg",
		Defense:    "sdfg",
		SpAttack:   "fdsg",
		SpDefense:  "sdfg",
		Speed:      "dfsg",
		Generation: "fdg",
		Legendary:  "fdsg",
	},
}

func TestPokemonService_GetPokemonByName(t *testing.T) {

	tests := []struct {
		name        string
		pokemonName string
		want        *model.Pokemon
		expectedErr error
	}{
		{
			name:        "prueba",
			pokemonName: "bulbasor",
			want: &model.Pokemon{
				Id:         "",
				Name:       "bulbasor",
				Type1:      "",
				Type2:      "",
				Total:      "",
				HP:         "",
				Attack:     "",
				Defense:    "",
				SpAttack:   "",
				SpDefense:  "",
				Speed:      "",
				Generation: "",
				Legendary:  "",
			},
			expectedErr: nil,
		},
		{
			name:        "pokemon not found ",
			pokemonName: "bulbassadivuh",
			expectedErr: errors.New("Pokemon not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &mocks.ServiceInterface{}

			ps.On("getPokemons").Return(prueba)

			ps.On("GetPokemonByName", tt.pokemonName).Return(tt.want, tt.expectedErr)

			got, err := ps.GetPokemonByName(tt.pokemonName)
			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr.Error(), err.Error())
			}
			assert.Equal(t, tt.want, got)

			if got != nil {
				assert.Equal(t, tt.pokemonName, got.Name)
			}

		})
	}
}

func TestPokemonService_GetPokemonByID(t *testing.T) {

	tests := []struct {
		ID          string
		pokemonName string
		want        *model.Pokemon
		expectedErr error
	}{
		{
			ID:          "1",
			pokemonName: "bulbasor",
			want: &model.Pokemon{
				Id:         "1",
				Name:       "",
				Type1:      "",
				Type2:      "",
				Total:      "",
				HP:         "",
				Attack:     "",
				Defense:    "",
				SpAttack:   "",
				SpDefense:  "",
				Speed:      "",
				Generation: "",
				Legendary:  "",
			},
			expectedErr: nil,
		},
		{
			ID:          "1231",
			pokemonName: "bulbassadivuh",
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.pokemonName, func(t *testing.T) {
			ps := &mocks.ServiceInterface{}

			ps.On("GetPokemonByID", tt.ID).Return(tt.want, tt.expectedErr)

			got, err := ps.GetPokemonByID(tt.ID)
			if tt.expectedErr != nil {
				fmt.Print(err.Error())
				fmt.Print(tt.expectedErr.Error())
				assert.Equal(t, tt.expectedErr.Error(), err.Error())
			}
			assert.Equal(t, tt.want, got)

			if got != nil {
				assert.Equal(t, tt.ID, got.Id)
			}
		})
	}
}

func TestPokemonService_getPokemons(t *testing.T) {
	type fields struct {
		servicePath string
	}
	tests := []struct {
		name   string
		fields fields
		want   []model.Pokemon
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PokemonService{
				servicePath: tt.fields.servicePath,
			}
			if got := p.getPokemons(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PokemonService.getPokemons() = %v, want %v", got, tt.want)
			}
		})
	}
}
