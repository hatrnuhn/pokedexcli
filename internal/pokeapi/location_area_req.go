package pokeapi

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullUrl := baseURL + endpoint
	if pageURL != nil {
		fullUrl = *pageURL
	}

	locationAreasResp := LocationAreasResp{}
	err := c.sendGetReq(fullUrl, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}
	return locationAreasResp, nil
}
