package repl

import "fmt"

func Pokedex(cfg *Config, args ...string) error {
	if len(cfg.Pokedex) == 0 {
		return fmt.Errorf("No tienes ningun pokemon! Empieza a atraparlos con el comando \"catch\"")
	}
	for _, pokemon := range cfg.Pokedex {
		fmt.Printf("	- %s\n", pokemon.Name)
	}
	return nil
}
