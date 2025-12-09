// Package main содержит точку входа основного приложения vhagar.
package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

const (
	defaultPort = 8080
)

type Config struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Port    int    `yaml:"port"`
}

func main() {
	config := Config{
		Name:    "vhagar",
		Version: "1.0.0",
		Port:    defaultPort,
	}

	data, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatalf("error marshaling YAML: %v", err)
	}

	fmt.Println("YAML output:")
	fmt.Println(string(data))

	var decoded Config
	err = yaml.Unmarshal(data, &decoded)
	if err != nil {
		log.Fatalf("error unmarshaling YAML: %v", err)
	}

	fmt.Printf("\nDecoded config: %+v\n", decoded)
}
