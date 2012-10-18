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
	default:
	    fmt.Fprintf(w, "hello")

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

func api_convertToJson(w http.ResponseWriter, qsoList []qslObject) {
    qsosJson, err := json.Marshal(qsoList)
    if err == nil {
	fmt.Fprintf(w, string(qsosJson))
    }
}
