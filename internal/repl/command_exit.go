package repl

import "os"

func Exit(cfg *Config, args ...string) error {
	os.Exit(0)
	return nil
}
