package repl

import (
	"errors"
	"fmt"
)

func Inspect(cfg *Config, args ...string) error {
	if len(args) == 0 {
		return errors.New("Debes especificar el nombre del Pokemon que deseas inspeccionar")
	}

	pok, ok := cfg.Pokedex[args[0]]
	if !ok {
		return errors.New("No has atrapado a ese pokemon todavia!")
	}

	fmt.Printf("Nombre: %s\n", pok.Name)
	fmt.Printf("Experiencia: %d\n", pok.BaseExperience)
	fmt.Printf("Peso: %d\n", pok.Weight)
	fmt.Printf("Altura: %d\n", pok.Height)
	fmt.Printf("Estadisticas:\n")
	for _, s := range pok.Stats {
		fmt.Printf("	- %s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Printf("Tipos:\n")
	for _, t := range pok.Types {
		fmt.Printf("	- %s\n", t.Type.Name)

	}
	return nil
}
