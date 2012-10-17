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
)

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<title>K2GXT QSL Cards</title>")
    fmt.Fprintf(w, "<style> h1 { text-align:center;}</style><h1><img src=\"../images/k2logo.jpg\" width=480 height=120></img></br>QSL Cards</h1></br>")
    for i := 1; i < len(jsontype); i++ {
	var call string = jsontype[i].Callsign
	fmt.Fprintf(w, "<a href=/view/"+call +">" + call + "</a></br>")
    }
}

func displayCard(w http.ResponseWriter, r *http.Request) {
    var callToCheck string = r.URL.Path[6:]
    fmt.Fprintf(w, "<title>" + callToCheck + "</title>")
    for i := 1; i < len(jsontype); i++ {
	if jsontype[i].Callsign == callToCheck {
	    fmt.Fprintf(w, "<h1>" + jsontype[i].Callsign + "</h1></br>")
	    fmt.Fprintf(w, "Date: " + jsontype[i].Date)
	    fmt.Fprintf(w, "</br>Mode: " + jsontype[i].Mode)
	    fmt.Fprintf(w, "</br>Frequency: " + jsontype[i].Frequency)
	    var fileName string = "../convertedCards/" + jsontype[i].Front_image
	    fmt.Fprintf(w, "</br></br>Front:</br><img src=\"" + fileName + "\" width=480 height=320 ></img>")
	    fmt.Fprintf(w, "</br><a href=../cards/" + jsontype[i].Front_image + "> Download full sized image </a>")
	    var backName string = "../../convertedCards/" + jsontype[i].Back_image
	    fmt.Fprintf(w, "</br></br>Back:</br><img src=\"" + backName + "\" width=480 height=320></img>")
	    fmt.Fprintf(w, "</br><a href=../../cards/" + jsontype[i].Back_image + "> Download full sized image </a>")
	}
    }
}
