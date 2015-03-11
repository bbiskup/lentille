package main

// Parsing of configuration file

import (
	//"github.com/kylelemons/go-gypsy/yaml"
	"log"
)

type Config interface {
}

func NewConfig(configFileName string) *Config {
	log.Printf("Using config file %s", configFileName)
	// result := new(config)
	return nil
}
