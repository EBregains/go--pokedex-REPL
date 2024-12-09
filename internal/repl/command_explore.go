package repl

import (
	"errors"
	"fmt"
)

func Explore(cfg *Config, args ...string) error {
	if len(args) == 0 {
		return errors.New("Debes ingresar el nombre de una ciudad para poder usar el comando \"explore\".")
	}
	locationResp, err := cfg.PokeapiClient.GetLocation(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Analizando la ciudad %s...\n", locationResp.Name)
	fmt.Println("Se encontraron los siguientes pokemons:")
	for i, pok := range locationResp.PokemonEncounters {
		fmt.Printf("%d. %s\n", i+1, pok.Pokemon.Name)
	}

	return nil
}
