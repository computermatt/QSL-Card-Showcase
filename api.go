/*
    QSL Card Showcase
    Copyright (C) 2012 Rochester Institute of Technology Amateur Radio Club, K2GXT 

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func apiGetCall(w http.ResponseWriter, r *http.Request) {
	var function = r.URL.Path[5:]
	switch true {
	case strings.Contains(function, "qso/callsign"):
		api_Callsign(w, r.URL.Path[18:])
	case strings.Contains(function, "qso/year"):
		api_Year(w, r.URL.Path[14:])
	case strings.Contains(function, "qso/mode"):
		api_Mode(w, r.URL.Path[14:])
	case strings.Contains(function, "qso/fullImage"):
		api_FullImage(w, r, r.URL.Path[19:])
	case strings.Contains(function, "qso/compressedImage"):
		api_CompressedImage(w, r, r.URL.Path[25:])
	case strings.Contains(function, "qso/thumbnailImage"):
		api_ThumbnailImage(w, r, r.URL.Path[24:])
	case strings.Contains(function, "qso/inBand"):
		api_Band(w, r.URL.Path[16:])
	case strings.Contains(function, "qso/country"):
		api_Country(w, r.URL.Path[17:])
	default:
		fmt.Fprintf(w, "error")
	}
}

func api_Callsign(w http.ResponseWriter, call string) {
	call = strings.ToUpper(call)
	qsoList := make([]qslObject, 0)
	foundCall := false
	for i := 0; i < len(qsls); i++ {
		if qsls[i].Callsign == call {
			foundCall = true
			qsoList = append(qsoList, qsls[i])
		}

	}
	if foundCall {
		api_convertToJson(w, qsoList)
	} else {
		fmt.Fprintf(w, "error")
	}
}

func api_Year(w http.ResponseWriter, year string) {
	qsoList := make([]qslObject, 0)
	foundContacts := false
	for i := 0; i < len(qsls); i++ {
		if strings.Contains(strings.ToUpper(qsls[i].Date), strings.ToUpper(year)) {
			foundContacts = true
			qsoList = append(qsoList, qsls[i])
		}
	}
	if foundContacts {
		api_convertToJson(w, qsoList)
	} else {
		fmt.Fprintf(w, "error")
	}
}

func api_Mode(w http.ResponseWriter, mode string) {
	qsoList := make([]qslObject, 0)
	foundContacts := false
	for i := 0; i < len(qsls); i++ {
		if strings.Contains(strings.ToUpper(qsls[i].Mode), strings.ToUpper(mode)) {
			foundContacts = true
			qsoList = append(qsoList, qsls[i])
		}
	}
	if foundContacts {
		api_convertToJson(w, qsoList)
	} else {
		fmt.Fprintf(w, "error")
	}
}

func api_FullImage(w http.ResponseWriter, r *http.Request, imageName string) {
	http.ServeFile(w, r, cardsFolder+imageName+fullType)
}

func api_CompressedImage(w http.ResponseWriter, r *http.Request, imageName string) {
	http.ServeFile(w, r, convertedFolder+imageName+convertedType)
}

func api_ThumbnailImage(w http.ResponseWriter, r *http.Request, imageName string) {
	http.ServeFile(w, r, resizedFolder+imageName+convertedType)
}

func api_Band(w http.ResponseWriter, band string) {
	qsoList := make([]qslObject, 0)
	foundContacts := false
	for i := 0; i < len(qsls); i++ {
		freq, err := strconv.ParseFloat(qsls[i].Frequency, 64)
		if (freq <= bandPlan[band].upper) && (freq >= bandPlan[band].lower) && (err == nil) {
			foundContacts = true
			qsoList = append(qsoList, qsls[i])
		}
	}
	if foundContacts {
		api_convertToJson(w, qsoList)
	} else {
		fmt.Fprintf(w, "error")
	}
}
func api_Country(w http.ResponseWriter, country string) {
	api_convertToJson(w, listOfContactsPerCountry[strings.ToUpper(country)])
}

func api_convertToJson(w http.ResponseWriter, qsoList []qslObject) {
	qsosJson, err := json.Marshal(qsoList)
	if err == nil {
		fmt.Fprintf(w, string(qsosJson))
	}
}
