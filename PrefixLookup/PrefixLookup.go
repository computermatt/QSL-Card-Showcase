/*
    PrefixLookup
    Copyright (C) 2012 Rochester Institute of Technology Amateur Radio Club, K2GXT 

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package PrefixLookup

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type QslObject struct {
	Callsign    string
	Date        string
	Front_image string
	Back_image  string
	Mode        string
	Frequency   string
}

var (
	prefixes map[string]interface{}
	country  string
)

func LoadPrefixes() {
	file, e := ioutil.ReadFile("PrefixLookup/prefixes.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	json.Unmarshal(file, &prefixes)
}

func CountryForCallsign(callsign string) string {
	country = ""
	callsign = strings.ToUpper(callsign)
	var currentMap = make(map[string]interface{})
	currentMap = prefixes

forLoop:
	for _, character := range callsign {

		switch cMap := currentMap[string(character)].(type) {
		case string:
			country = cMap
			break forLoop
		case nil:
			if currentMap["*"] != nil {
				country = currentMap["*"].(string)

			} else {
				return "ERROR"
			}
			break forLoop
		default:
			currentMap = cMap.(map[string]interface{})
		}
	}

	return country
}

func ListCountries(qsls []QslObject) []string {
	listOfCountries := make([]string, len(qsls))

	for i := 0; i < len(qsls); i++ {
		listOfCountries[i] = CountryForCallsign(qsls[i].Callsign)
	}
	removeDuplicates(&listOfCountries)
	sort.Strings(listOfCountries)
	return listOfCountries

}

func removeDuplicates(listOfCountries *[]string) {
	found := make(map[string]bool)
	j := 0
	for i, x := range *listOfCountries {
		if !found[x] {
			found[x] = true
			(*listOfCountries)[j] = (*listOfCountries)[i]
			j++
		}
	}
	*listOfCountries = (*listOfCountries)[:j]
}
