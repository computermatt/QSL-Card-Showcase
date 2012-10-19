/*
    QSL Card Showcase
    Copyright (C) 2012 Rochester Institute of Technology Amateur Radio Club, K2GXT 

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package main

import (
    "fmt"
    "net/http"
    "strings"
    "encoding/json"
    "strconv"
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
    http.ServeFile(w, r, cardsFolder + imageName + fullType)
}

func api_CompressedImage(w http.ResponseWriter, r *http.Request, imageName string) {
    http.ServeFile(w, r, convertedFolder + imageName + convertedType)
}

func api_ThumbnailImage(w http.ResponseWriter, r *http.Request, imageName string) {
    http.ServeFile(w, r, resizedFolder + imageName + convertedType)
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

func api_convertToJson(w http.ResponseWriter, qsoList []qslObject) {
    qsosJson, err := json.Marshal(qsoList)
    if err == nil {
	fmt.Fprintf(w, string(qsosJson))
    }
}
