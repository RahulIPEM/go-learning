package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

// Config holds the configuration
type Config struct {
	// configuration fields goes here.
}

func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "Can't open configuration file.")
	}

	defer file.Close()

	cfg := &Config{}
	// Parse file here
	return cfg, nil
}

func setupLogging() {
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	log.SetOutput(out)
}

func main() {
	setupLogging()
	cfg, err := readConfig("path/to/config.toml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		log.Printf("Error:: %+v", err)
		os.Exit(1)
	}

	// Normal operation
	fmt.Println(cfg)
}
