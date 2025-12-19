package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/milkman-dev/pokedexcli/pokecache"
)

func (c *Client) GetLocations(pageURL *string) (Locations, error) {
	locationsResp := Locations{}
	url := baseURL + "location-area"
	if pageURL != nil {
		url = *pageURL
	}

	cache := pokecache.NewCache(10 * time.Second)
	if v, ok := cache.Get(url); ok {
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
	cache.Add(url, data)

	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return Locations{}, err
	}

	return locationsResp, nil
}
