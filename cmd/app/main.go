package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
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
		Port:    8080,
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
