package pokeapi

func (c *Client) ListPokemonsInArea(areaName string) (PokemonsInAreaResp, error) {
	endpoint := "/location-area/" + areaName
	fullUrl := baseURL + endpoint
	pokemonInAreaResp := PokemonsInAreaResp{}
	err := c.sendGetReq(fullUrl, &pokemonInAreaResp)
	if err != nil {
		return PokemonsInAreaResp{}, err
	}
	return pokemonInAreaResp, nil
}
