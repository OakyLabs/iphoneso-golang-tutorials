package main

import (
	"encoding/json"
	"log"
	"net/http"

	"pokemonso/pokemon"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello there"))
	})

	mux.HandleFunc("GET /random", func(w http.ResponseWriter, r *http.Request) {
		pokemonster, err := pokemon.GetRandomPokemon()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))

			return
		}

		bytes, err := json.Marshal(pokemonster)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))

			return
		}

		w.Write(bytes)
	})

	log.Fatal(http.ListenAndServe(":3000", mux))
}
