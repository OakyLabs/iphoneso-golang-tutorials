package pokemon

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

type Pokemon struct {
	Name   string         `json:"name"`
	Order  int            `json:"order"`
	Weigth int            `json:"weight"`
	Forms  []PokemonForms `json:"forms"`
}

type PokemonForms struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

var quoteUnquoteRandomPokemon = []string{"pikachu", "bulbasaur", "ditto", "raticate", "pidgeot", "charizard", "mew", "rattata"}

var quoteUnquoteRandomPokemonMap = map[string]Pokemon{}

// Pokemon, error
// Pokemon, nikl
// Pokemon, errors.New("Kaboom")

func GetRandomPokemon() (*Pokemon, error) {
	ranIndex := rand.Intn(len(quoteUnquoteRandomPokemon))

	ranPoke := quoteUnquoteRandomPokemon[ranIndex]

	pokemon, ok := quoteUnquoteRandomPokemonMap[ranPoke]

	if !ok {
		fmt.Println("Requesting new pokemon")
		url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", ranPoke)

		// POST -> nil
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("Error in request")
			return nil, err
		}

		client := &http.Client{}

		response, err := client.Do(request)

		if err != nil {
			fmt.Println("Error in response")
			return nil, err
		}

		poke := Pokemon{}

		err = json.NewDecoder(response.Body).Decode(&poke)

		if err != nil {
			fmt.Println("Error in unmarshallingh", err.Error())
			return nil, err
		}

		quoteUnquoteRandomPokemonMap[ranPoke] = poke
		// http calls
		return &poke, nil
	}

	return &pokemon, nil
}
