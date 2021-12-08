package controller

import (
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/rayl17/academy-go-q42021/mocks"
	"github.com/rayl17/academy-go-q42021/model"
	"github.com/rayl17/academy-go-q42021/usecase"
	"github.com/stretchr/testify/assert"
)

func TestPokemonController_GetPokemonByName(t *testing.T) {
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
			c := &mocks.ServiceInterface{}
			c.On("GetPokemonByName", tt.pokemonName).Return(tt.want, tt.expectedErr)

			got, err := c.GetPokemonByName(tt.pokemonName)
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

func TestPokemonController_GetPokemonByID(t *testing.T) {
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
			c := &mocks.ServiceInterface{}
			c.On("GetPokemonByID", tt.ID).Return(tt.want, tt.expectedErr)

			got, err := c.GetPokemonByID(tt.ID)
			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr.Error(), err.Error())
			}
			assert.Equal(t, tt.want, got)

			if got != nil {
				assert.Equal(t, tt.ID, got.Id)
			}
		})
	}
}

func TestPokemonController_PostPokemon(t *testing.T) {
	type fields struct {
		usecase usecase.Usecaseinterface
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &PokemonController{
				usecase: tt.fields.usecase,
			}
			c.PostPokemon(tt.args.w, tt.args.r)
		})
	}
}
