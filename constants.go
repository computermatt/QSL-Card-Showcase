package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	qsoFile         string
	convertedFolder string
	cardsFolder     string
	resizedFolder   string
	imagesFolder    string
	callsign        string
	logoFileName    string
	convertedType   string
	fullType        string
)

func parseConstants() {
	file, e := ioutil.ReadFile("settings.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	constants := make(map[string]string)
	json.Unmarshal(file, &constants)
	qsoFile = constants["qsoFile"]
	convertedFolder = constants["convertedFolder"]
	cardsFolder = constants["cardsFolder"]
	resizedFolder = constants["resizedFolder"]
	imagesFolder = constants["imagesFolder"]
	callsign = constants["callsign"]
	logoFileName = constants["logoFileName"]
	convertedType = constants["convertedType"]
	fullType = constants["fullType"]

}
