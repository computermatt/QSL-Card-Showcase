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
    fmt.Fprintf(w, "<title>" + callsign + " QSL Cards</title>")
    fmt.Fprintf(w, "<style> h1 { text-align:center;}</style><h1><img src=\"" + imagesFolder + logoFileName + "\" width=480 height=120></img></br>QSL Cards</h1></br>")
    for i := 0; i < len(qsls); i++ {
	var call string = qsls[i].Callsign
	fmt.Fprintf(w, "<a href=/view/"+call +">" + call + "</a>  <img src=\"../resizedCards/" + qsls[i].Front_image + convertedType + "\" width=100 height=60></img>  <img src=\"../resizedCards/" + qsls[i].Back_image + convertedType + "\" width=100 height=60></img></br>")
    }
}

func displayCard(w http.ResponseWriter, r *http.Request) {
    var callToCheck string = r.URL.Path[6:]
    fmt.Fprintf(w, "<title>" + callToCheck + "</title>")
    fmt.Fprintf(w, "<img src=\"" + imagesFolder + logoFileName + "\" width=320 height=80></img>")
    for i := 0; i < len(qsls); i++ {
	if qsls[i].Callsign == callToCheck {
	    fmt.Fprintf(w, "<h1>" + qsls[i].Callsign + "</h1></br>")
	    fmt.Fprintf(w, "Date: " + qsls[i].Date)
	    fmt.Fprintf(w, "</br>Mode: " + qsls[i].Mode)
	    fmt.Fprintf(w, "</br>Frequency: " + qsls[i].Frequency)
	    var fileName string = "../convertedCards/" + qsls[i].Front_image
	    fmt.Fprintf(w, "</br></br>Front:</br><img src=\"" + fileName + convertedType + "\" width=480 height=320 ></img>")
	    fmt.Fprintf(w, "</br><a href=../cards/" + qsls[i].Front_image + fullType + "> Download full sized image </a>")
	    var backName string = "../../convertedCards/" + qsls[i].Back_image
	    fmt.Fprintf(w, "</br></br>Back:</br><img src=\"" + backName + convertedType + "\" width=480 height=320></img>")
	    fmt.Fprintf(w, "</br><a href=../../cards/" + qsls[i].Back_image + fullType + "> Download full sized image </a>")
	}
    }
}
