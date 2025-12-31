package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(pokemonURL *string) (Pokemon, error) {
	url := baseURL + "pokemon/" + *pokemonURL
	pokemon := Pokemon{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err  != nil {
		return Pokemon{}, err
	}

	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, nil
	}

	return pokemon, nil
}