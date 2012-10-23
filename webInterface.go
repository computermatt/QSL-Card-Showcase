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
    "strconv"
)

func index(w http.ResponseWriter, r *http.Request) {
    var startCard , _  = strconv.ParseInt(r.URL.Path[1:], 0, 0)
    fmt.Fprintf(w, "<title>" + callsign + " QSL Cards</title>")
    fmt.Fprintf(w, "<div style=\"text-align:center\"><h1><img src=\"" + imagesFolder + logoFileName + "\" width=480 height=120></img></br>QSL Cards</h1></br></div>")
    var endCard int = int(startCard) + 20
    if endCard > len(qsls) {
	endCard = len(qsls)
    }
    var numInRow int = 0
    fmt.Fprintf(w, "<table boarder=\"1\" align=\"center\"><tr>")
    for i := int(startCard); i < endCard; i++ {
	var call string = qsls[i].Callsign
	fmt.Fprintf(w, "<td><div style=\"text-align:center\"><a href=/view/"+ call +">" + call + "</a></br><img src=\"" + resizedFolder + qsls[i].Front_image + convertedType + "\" width=100 height=60></img>  <img src=\"" + resizedFolder + qsls[i].Back_image + convertedType + "\" width=100 height=60></img></div><p>     </p></td>" )
	if numInRow == 3 {
	    fmt.Fprintf(w, "</tr><tr>")
	    numInRow = 0
	} else {
	    numInRow++
	}
    }
    var newStart int = int(startCard) - 20
    if newStart <= 0 {
	newStart = 0
    }
    fmt.Fprintf(w, "<tr><td><div style=\"text-align:left\"><a href=/" + strconv.Itoa(newStart) + ">Back</a></div></td><td></td><td></td><td>  <div style=\"text-align:right\"> <a href=/" + strconv.Itoa(endCard) + ">Next</a></td></tr></table>")
}

func displayCard(w http.ResponseWriter, r *http.Request) {
    var callToCheck string = r.URL.Path[6:]
    fmt.Fprintf(w, "<title>" + callToCheck + "</title>")
    fmt.Fprintf(w, "<div style=\"text-align:center\"><img src=\"" + imagesFolder + logoFileName + "\" width=320 height=80></img>")
    for i := 0; i < len(qsls); i++ {
	if qsls[i].Callsign == callToCheck {
	    fmt.Fprintf(w, "<u><h1>" + qsls[i].Callsign + "</h1></u>")
	    fmt.Fprintf(w, "<b>Date: " + qsls[i].Date)
	    fmt.Fprintf(w, "</br>Mode: " + qsls[i].Mode)
	    fmt.Fprintf(w, "</br>Frequency: " + qsls[i].Frequency + "</b>")
	    var fileName string = convertedFolder+ qsls[i].Front_image
	    fmt.Fprintf(w, "</br><table boarder=\"1\" align=\"center\"><tr><td><img src=\"" + fileName + convertedType + "\" width=480 height=320 ></img>")
	    fmt.Fprintf(w, "</br><div style=\"text-align:center\"><a href=" + cardsFolder + qsls[i].Front_image + fullType + "> Download full sized image </a></div></td><td>")
	    var backName string = "../" + convertedFolder + qsls[i].Back_image
	    fmt.Fprintf(w, "<img src=\"" + backName + convertedType + "\" width=480 height=320></img>")
	    fmt.Fprintf(w, "</br><div style=\"text-align:center\"><a href=../" + cardsFolder + qsls[i].Back_image + fullType + "> Download full sized image </a></div></td></tr>")
	    fmt.Fprintf(w, "</table></br>")
	}
    }
}
