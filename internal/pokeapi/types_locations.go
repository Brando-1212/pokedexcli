package pokeapi

//RespShallowLocations
type RespShallowLocations struct {
	Count    int     `"json:"count"`
	Next     *string `json:"previous"`
	Previous *string `json: "previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}`json:"results"`
}