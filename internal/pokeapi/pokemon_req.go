package pokeapi

import "fmt"

// this should return one certain pokemon
func (c *Client) PokemonReq(pokemon string) (PokemonResp, error) {
	endpoint := "/pokemon/" + pokemon
	fullUrl := baseURL + endpoint

	pokemonResp := PokemonResp{}

	err := c.sendGetReq(fullUrl, &pokemonResp)
	if err != nil {
		return PokemonResp{}, err
	}

	return pokemonResp, nil
}

func (c *Client) PokemonsListReq() (PokemonsListResp, error) {
	endpoint := fmt.Sprintf("/pokemon?limit=%v", "1")
	fullUrl := baseURL + endpoint

	pokemonsListResp := PokemonsListResp{}
	err := c.sendGetReq(fullUrl, &pokemonsListResp)
	if err != nil {
		return PokemonsListResp{}, err
	}

	pokemonsCount := pokemonsListResp.Count
	fullUrl = baseURL + fmt.Sprintf("/pokemon?limit=%v", pokemonsCount)
	err = c.sendGetReq(fullUrl, &pokemonsListResp)
	if err != nil {
		return PokemonsListResp{}, err
	}

	return pokemonsListResp, nil
}
