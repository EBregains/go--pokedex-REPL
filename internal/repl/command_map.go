package repl

import (
	"fmt"
)

func Map(cfg *Config, args ...string) error {
	locationsResp, err := cfg.PokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous
	cfg.LocationsPage++

	for i, loc := range locationsResp.Results {
		fmt.Printf("%d. %s\n", (i+1)*cfg.LocationsPage, loc.Name)
	}
	return nil
}

func Mapb(cfg *Config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		cfg.LocationsPage = 0
		return fmt.Errorf("You are on the first page")
	}

	locationsResp, err := cfg.PokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous
	cfg.LocationsPage--

	for i, loc := range locationsResp.Results {
		fmt.Printf("%d. %s\n", (i+1)*cfg.LocationsPage, loc.Name)
	}
	return nil
}
