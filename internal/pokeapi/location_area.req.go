package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullUrl := baseURL + endpoint
	if pageURL != nil {
		fullUrl = *pageURL
	}

	dat, ok := c.cache.Get(fullUrl)
	if ok {
		// unmarshal dat and store it in pointer to locationAreasResp
		fmt.Println("cache hit!")
		locationAreasResp := LocationAreasResp{}
		err := json.Unmarshal(dat, &locationAreasResp)
		if err != nil {
			return LocationAreasResp{}, err
		}

		return locationAreasResp, nil
	}
	fmt.Println("cache miss!")

	// create new GET request, since it's a GET, we send nil body
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	// execute the request, get HTTP response or error
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}

	// close response when exiting function
	defer resp.Body.Close()

	// StatusCode > 399 means a client-side error
	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	// unmarshal dat and store it in pointer to locationAreasResp
	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(fullUrl, dat)

	return locationAreasResp, nil
}
