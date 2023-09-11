package main

import (
	"fmt"
	"os"
)

// create folders for mods
func createTempFiles() {
	for _, mod := range config.Mods {
		err := os.MkdirAll(fmt.Sprintf("mods/%s", mod), 0755)
		if err != nil {
			logg.Fatal(err)
		}
	}
}
