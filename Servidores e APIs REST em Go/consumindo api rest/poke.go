package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// response
type Resposta struct{
	Nome string `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}


type Pokemon struct {
	Numero int `json:"entry_number"`
	Especie PokemonSpecies `json:"pokemon_species"`
}
type PokemonSpecies struct {
	Nome string `json:"name"`
	Url string `json:"url"`
}

func main() {
	resposta, err := http.Get("https://pokeapi.co/api/v2/pokedex/kanto")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	dadosResposta, err := io.ReadAll(resposta.Body)
	if err != nil {
		log.Println(err)
	}

	var objetoResposta Resposta
	json.Unmarshal(dadosResposta, &objetoResposta)

	fmt.Println(objetoResposta.Nome)
	fmt.Println(objetoResposta.Pokemon)

	for _, pokemon := range objetoResposta.Pokemon {
		fmt.Println(pokemon.Especie.Nome)
	}
}