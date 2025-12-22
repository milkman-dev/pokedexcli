package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemons(locationUrl *string) (AreaPokemons, error) {
	url := baseURL + "location-area/" + *locationUrl

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AreaPokemons{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return AreaPokemons{}, err
	}
	defer resp.Body.Close()

	respAreaPokemons := AreaPokemons{}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return AreaPokemons{}, err
	}

	if err := json.Unmarshal(data, &respAreaPokemons); err != nil {
		return AreaPokemons{}, err
	}

	return respAreaPokemons, nil
}
