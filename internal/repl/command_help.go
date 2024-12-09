package repl

import "fmt"

func Help(cfg *Config, args ...string) error {
	fmt.Println()
	fmt.Println("Bienvenido a tu Pokedex!")
	fmt.Println("Los comandos disponibles son:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
