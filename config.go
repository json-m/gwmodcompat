package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var config Config

type Config struct {
	Mods []string `yaml:"mods"`
}

func init() {
	err := readConfig()
	if err != nil {
		log.Fatal(err)
	}
}

// read config file
func readConfig() error {
	// open config file
	f, err := os.Open("config.yml")
	if err != nil {
		log.Fatal("readConfig.open:", err)
	}

	// read config file
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("readConfig.Decode:", err)
	}

	f.Close()
	return nil
}
