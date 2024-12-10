package main

import (
	"time"

	"github.com/EBregains/REPL-Pokedex/internal/pokeapi"
	"github.com/EBregains/REPL-Pokedex/internal/repl"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &repl.Config{
		PokeapiClient: pokeClient,
		Pokedex:       make(map[string]pokeapi.Pokemon),
	}
	repl.Start(cfg)
}
