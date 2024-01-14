package pokeapi

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
