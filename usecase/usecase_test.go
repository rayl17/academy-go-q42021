package usecase

import (
	"errors"
	"reflect"
	"testing"

	"github.com/rayl17/academy-go-q42021/mocks"
	"github.com/rayl17/academy-go-q42021/model"
	"github.com/stretchr/testify/assert"
)

func TestPokemonUsecase_GetPokemonByName(t *testing.T) {
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
			u := &mocks.Usecaseinterface{}

			u.On("GetPokemonByName", tt.pokemonName).Return(tt.want, tt.expectedErr)
			got, err := u.GetPokemonByName(tt.pokemonName)
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

func TestPokemonUsecase_GetPokemonByID(t *testing.T) {
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
			u := &mocks.Usecaseinterface{}
			u.On("GetPokemonByID", tt.ID).Return(tt.want, tt.expectedErr)

			got, err := u.GetPokemonByID(tt.ID)
			if tt.expectedErr != nil {
				t.Errorf("PokemonUsecase.GetPokemonByID() error = %v, wantErr %v", err, tt.expectedErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PokemonUsecase.GetPokemonByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPokemonUsecase_SavePokemonByNameAPI(t *testing.T) {

	tests := []struct {
		name        string
		want        string
		expectedErr error
		pokemonName string
	}{
		{
			name:        "prueba",
			want:        "Saved succesfully",
			expectedErr: nil,
			pokemonName: "bulbasor",
		},
		{
			name:        "pokemon not found ",
			pokemonName: "bulbassadivuh",
			expectedErr: errors.New("Pokemon not found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &mocks.Usecaseinterface{}

			u.On("SavePokemonByNameAPI", tt.pokemonName).Return(tt.want, tt.expectedErr)
			got, err := u.SavePokemonByNameAPI(tt.pokemonName)
			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr.Error(), err.Error())
				return
			}
			if got != tt.want {
				t.Errorf("PokemonUsecase.SavePokemonByNameAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}
