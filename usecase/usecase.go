package usecase

///cambir por use case
import (
	"errors"

	"github.com/rayl17/academy-go-q42021/model"
	service "github.com/rayl17/academy-go-q42021/service/csv"
)

type Usecaseinterface interface {
	GetPokemonByName(name string) (*model.Pokemon, error)
	GetPokemonByID(id string) (*model.Pokemon, error)
	SavePokemonByNameAPI(name string) (string, error)
}

type PokemonUsecase struct {
	service service.ServiceInterface
}

func NewUseCase(ps service.ServiceInterface) *PokemonUsecase {

	return &PokemonUsecase{
		service: ps,
	}

}

func (u *PokemonUsecase) GetPokemonByName(name string) (*model.Pokemon, error) {

	pokemon, err := u.service.GetPokemonByName(name)

	if err != nil {

		return nil, errors.New("Pokemon not found")
	}

	return pokemon, nil

}

// int
func (u *PokemonUsecase) GetPokemonByID(id string) (*model.Pokemon, error) {

	pokemon, err := u.service.GetPokemonByID(id)

	if err != nil {

		return nil, errors.New("Pokemon not found")
	}

	return pokemon, nil
}

func (u *PokemonUsecase) SavePokemonByNameAPI(name string) (string, error) {

	message, err := u.service.SearchPokemonApi(name)

	if err != nil {

		return "", errors.New("Pokemon not found")
	}

	return message, nil
}