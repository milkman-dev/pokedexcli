package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocations(pageURL *string) (Locations, error) {
	locationsResp := Locations{}
	url := baseURL + "location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if v, ok := c.cache.Get(url); ok {
		fmt.Println("Using cache...")
		if err := json.Unmarshal(v, &locationsResp); err != nil {
			return Locations{}, err
		} else {

			return locationsResp, nil
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Locations{}, err
	}
	c.cache.Add(url, data)

	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return Locations{}, err
	}

	return locationsResp, nil
}
