package main

import (
	"gwmodcompat/pkg"
	"log"
	"os"
	"time"
)

var logg *log.Logger

func init() {
	// create own logger
	logg = log.New(os.Stdout, "*** ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	logg.Println("modcompat server startup")

	// create work/temp files
	createTempFiles()

	// background tasks
	go watcher()
	startHttpServer()
}

// 10min loop to check for updates
func watcher() {
	for {
		for _, mod := range config.Mods {
			logg.Println("checking", mod)
			pkg.Download(mod)
		}
		time.Sleep(10 * time.Minute)
	}
}
