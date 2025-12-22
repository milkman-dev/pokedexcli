package pokeapi

type Locations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Pokemon struct {
	Name string `json:"name"`
}

type AreaPokemons struct {
	PokemonEncounters []struct {
		Pokemon Pokemon
	} `json:"pokemon_encounters"`
}
