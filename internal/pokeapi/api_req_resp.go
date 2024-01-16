package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) sendGetReq(fullUrl string, v interface{}) error {
	dat, ok := c.Cache.Get(fullUrl)
	if ok {
		fmt.Println("cache hit!")
		err := json.Unmarshal(dat, v)
		if err != nil {
			return err
		}
		return nil
	}
	fmt.Println("cache miss!")

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(dat, v)
	if err != nil {
		return err
	}

	c.Cache.Add(fullUrl, dat)

	return nil
}
