package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/EBregains/REPL-Pokedex/internal/pokeapi"
)

type Config struct {
	PokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	LocationsPage    int
	Pokedex          map[string]pokeapi.Pokemon
}

func Start(cfg *Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		command := words[0]
		args := []string{}

		if len(words) > 1 {
			args = words[1:]
		}

		cmd, exists := getCommands()[command]
		if !exists {
			fmt.Printf("%s: Comando desconocido.\n", command)
			continue
		}

		err := cmd.handler(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}

	}
}

func cleanInput(text string) []string {
	textInLowerCase := strings.ToLower(text)
	words := strings.Fields(textInLowerCase)
	return words
}

type cliCommands struct {
	name        string
	description string
	handler     func(cfg *Config, args ...string) error
}

func getCommands() map[string]cliCommands {
	return map[string]cliCommands{
		"help": {
			name:        "help",
			description: "Muestra este menu y lista los comandos disponibles.",
			handler:     Help,
		},
		"exit": {
			name:        "exit",
			description: "Cierra el programa y finaliza su ejecucion.",
			handler:     Exit,
		},
		"map": {
			name:        "map",
			description: "Muestra las ubicaciones en el Mundo Pokemon, mostrando de a 20 ubicaciones por vez",
			handler:     Map,
		},
		"mapb": {
			name:        "mapb",
			description: "Muestra las 20 ubicaciones anteriores a las ultimas ubicaciones mostradas.",
			handler:     Mapb,
		},
		"explore": {
			name:        "explore",
			description: "Muestra los pokemones disponibles en el area seleccionada",
			handler:     Explore,
		},
		"catch": {
			name:        "catch",
			description: "Intenta cazar a un pokemon usando \"catch [nombre-pokemon]\"!",
			handler:     Catch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspecciona las stats de un Pokemon de tu Pokedex con \"inspect [nombre-pokemon]\"!",
			handler:     Inspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lista tu pokedex!",
			handler:     Pokedex,
		},
	}
}
