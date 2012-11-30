/*
    QSL Card Showcase
    Copyright (C) 2012 Rochester Institute of Technology Amateur Radio Club, K2GXT 

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package main

import (
	PrefixLookup "./PrefixLookup"
	"strings"
)

func getListOfCountries() []string {
	convertedQsls := make([]PrefixLookup.QslObject, len(qsls))
	for i := 0; i < len(qsls); i++ {
		qsl := qsls[i]
		var d PrefixLookup.QslObject
		d.Callsign = qsl.Callsign
		d.Back_image = qsl.Back_image
		d.Date = qsl.Date
		d.Frequency = qsl.Frequency
		d.Front_image = qsl.Front_image
		d.Mode = qsl.Mode
		convertedQsls[i] = d
	}
	return PrefixLookup.ListCountries(convertedQsls)

}

func getContactsPerCountry() map[string][]qslObject {
	tmpCountry := make(map[string][]qslObject)
	for i := 0; i < len(qsls); i++ {
		country := strings.ToUpper(PrefixLookup.CountryForCallsign(qsls[i].Callsign))
		tmpCountry[country] = append(tmpCountry[country], qsls[i])
	}
	return tmpCountry
}
