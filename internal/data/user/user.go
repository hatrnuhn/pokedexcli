package user

import (
	"errors"
	"fmt"

	"github.com/hatrnuhn/pokedexcli/internal/pokeapi"
)

type UserPokedex struct {
	userPokedex map[string]pokeapi.PokemonResp
}

func NewPokedex() UserPokedex {
	return UserPokedex{
		userPokedex: make(map[string]pokeapi.PokemonResp),
	}
}

func (pdx *UserPokedex) Add(pkName string, pokemon pokeapi.PokemonResp) error {
	_, ok := pdx.userPokedex[pkName]
	if ok {
		return fmt.Errorf("%v already exists in your Pokedex", pkName)
	}

	pdx.userPokedex[pkName] = pokemon
	return nil
}

func (pdx *UserPokedex) Remove(pkName string) error {
	_, ok := pdx.userPokedex[pkName]
	if !ok {
		return fmt.Errorf("%v doesn't exist in your Pokedex", pkName)
	}
	delete(pdx.userPokedex, pkName)
	return nil
}

func (pdx *UserPokedex) Get(pkName string) (pokeapi.PokemonResp, bool) {
	val, ok := pdx.userPokedex[pkName]
	return val, ok
}

func (pdx *UserPokedex) GetUserPokedex() (map[string]pokeapi.PokemonResp, error) {
	if len(pdx.userPokedex) == 0 {
		return pdx.userPokedex, errors.New("pokedex is empty")
	}

	return pdx.userPokedex, nil
}
