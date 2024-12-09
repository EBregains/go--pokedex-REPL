package repl

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"time"
)

func Catch(cfg *Config, args ...string) error {
	if len(args) == 0 {
		return errors.New("Debes seleccionar un Pokemon para poder usar el comando \"catch\"")
	}

	pokemon, err := cfg.PokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Lanzando una Pokebola a %s\n", pokemon.Name)
	time.Sleep(1000)
	result := 1.0 / float32(rand.IntN(pokemon.BaseExperience))
	if result > float32(0.02) {
		cfg.Pokedex[pokemon.Name] = pokemon
		fmt.Printf("Has atrapado a %s!\n", pokemon.Name)
		fmt.Println("Ahora puedes inspeccionar tu Pokemon con el comando inspect.")
	} else {
		fmt.Printf("%s ha escapado...\n", pokemon.Name)
	}
	return nil
}
