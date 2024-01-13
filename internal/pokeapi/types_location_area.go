package pokeapi

type LocationAreasResp struct {
	Count int `json:"count"`

	// pointer because next and prev may store null
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
