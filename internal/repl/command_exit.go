package repl

import (
	"fmt"
	"os"
)

func Exit(cfg *Config, args ...string) error {

	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
