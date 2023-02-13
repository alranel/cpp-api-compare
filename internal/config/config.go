package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type LibraryConfig struct {
	Name        string
	Description string
	Link        string
	Include     []string
	Exclude     []string
	Defines     []string
	Namespaces  []string
}

// Config represents a comparison job configuration
type Config struct {
	Libraries                      []LibraryConfig
	Classes, Singletons, Functions []string
}

func (c *Config) ReadFromFile(path string) {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Error reading config file: #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Error parsingn config file: %v", err)
	}
}
