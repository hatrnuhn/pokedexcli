package user

import (
	"testing"

	"github.com/hatrnuhn/pokedexcli/internal/pokeapi"
)

func TestCreateUserPokedex(t *testing.T) {
	pdx := NewPokedex()
	if pdx.userPokedex == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetUserPokedex(t *testing.T) {
	pdx := NewPokedex()

	cases := []struct {
		inputKey   string
		inputVal   pokeapi.PokemonResp
		expectedOk bool
	}{
		{
			inputKey: "key1",
			inputVal: pokeapi.PokemonResp{
				Name: "val1",
			},
			expectedOk: true,
		}, {
			inputKey: "key2",
			inputVal: pokeapi.PokemonResp{
				Name: "val2",
			},
			expectedOk: true,
		},
	}

	for _, cs := range cases {

		pdx.Add(cs.inputKey, cs.inputVal)

		actual, ok := pdx.Get(cs.inputKey)

		if !ok {
			t.Error("could not find ", cs.inputKey)
		}
		if string(actual.Name) != string(cs.inputVal.Name) {
			t.Errorf("value does not match: %v actual vs %v case", actual, cs.inputVal)
		}
	}
}
