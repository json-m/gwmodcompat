package main

import (
	"gwmodcompat/pkg"
	"log"
	"os"
)

var logg *log.Logger

func init() {
	// create own logger
	logg = log.New(os.Stdout, "*** ", log.Lshortfile)

}

func main() {
	logg.Println("modcompat server startup")

	// create work/temp files
	if _, err := os.Stat("mods"); os.IsNotExist(err) {
		logg.Println("creating work/temp files")
		createTempFiles()
	}

	//pkg.TestGithub()
	pkg.Download("knowthyenemy")

	startHttpServer()
}
