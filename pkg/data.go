package pkg

import (
	"encoding/json"
	"os"
)

func init() {
	err := readData()
	if err != nil {
		panic(err)
	}
}

// Data for each mod to track - saved to data.json
type Data struct {
	KnowThyEnemy KnowThyEnemy `json:"knowthyenemy"`
	ArcDps       ArcDps       `json:"arcdps"`
}

type History struct {
	Version  string `json:"version"`
	Checksum string `json:"checksum"`
}

var modData Data

// RetrieveData for external calls
func RetrieveData() Data {
	return modData
}

// placeholder
func CreateData() error {
	// create data file
	f, err := os.Create("data.json")
	if err != nil {
		panic(err)
	}
	f.Close()

	// write data file
	err = writeData()
	if err != nil {
		panic(err)
	}

	return nil
}

// writes Data to file
func writeData() error {
	// open data file
	f, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	// write data file
	encoder := json.NewEncoder(f)
	err = encoder.Encode(&modData)
	if err != nil {
		panic(err)
	}
	f.Close()

	return nil
}

// reads Data from file
func readData() error {
	// open data file
	f, err := os.Open("data.json")
	if err != nil {
		return err
	}

	// read data file
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&modData)
	if err != nil {
		return err
	}
	f.Close()

	return nil
}
