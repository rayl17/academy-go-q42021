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
	GetPokemonsConcurrently(types string, items int, items_per_worker int) (*[]model.Pokemon, error)
}

type PokemonUsecase struct {
	service service.ServiceInterface
}

func NewUseCase(ps service.ServiceInterface) *PokemonUsecase {
	return &PokemonUsecase{
		service: ps,
	}

}

/// in this metod we requierd an Name to find a pokemon Name to match the given name, returns the pokemon that matches , or an error
func (u *PokemonUsecase) GetPokemonByName(name string) (*model.Pokemon, error) {
	pokemon, err := u.service.GetPokemonByName(name)
	if err != nil {
		return nil, errors.New("Pokemon not found")
	}

	return pokemon, nil
}

/// in this metod we requierd an ID to find a pokemon ID to match the given ID, returns the pokemon that matches , or an error
func (u *PokemonUsecase) GetPokemonByID(id string) (*model.Pokemon, error) {
	pokemon, err := u.service.GetPokemonByID(id)
	if err != nil {
		return nil, errors.New("Pokemon not found")
	}

	return pokemon, nil
}

/// in this metod we requierd the pokemon name to save in the CSV file , if the method executes succesfully it returs succes if not returns an error
func (u *PokemonUsecase) SavePokemonByNameAPI(name string) (string, error) {
	message, err := u.service.SearchPokemonApi(name)
	if err != nil {
		return "", errors.New("Pokemon not found")
	}

	return message, nil
}

///In this method we require types string, items int, items_per_worker int) to get pokemons concurrently ,and we return an array of pokemons
// requires
func (u *PokemonUsecase) GetPokemonsConcurrently(types string, items int, items_per_worker int) (*[]model.Pokemon, error) {
	pokemon, err := u.service.GetPokemonsByTypes(types, items, items_per_worker)
	if err != nil {

		return nil, errors.New("Something went wrong")
	}

	return pokemon, nil
}
