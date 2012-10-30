/*
    QSL Card Showcase
    Copyright (C) 2012 Rochester Institute of Technology Amateur Radio Club, K2GXT 

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package main

import (
    "fmt"
    "net/http"
    "strconv"
    "strings"
)

func index(w http.ResponseWriter, r *http.Request) {
    var startCard , _  = strconv.ParseInt(r.URL.Path[1:], 0, 0)
    fmt.Fprintf(w, "<title>" + callsign + " QSL Cards</title>")
    fmt.Fprintf(w, "<div style=\"text-align:center\"><h1><a href=/><img src=\"" + imagesFolder + logoFileName + "\" width=480 height=120></a></img></br>QSL Cards</h1></br></div>")
    var endCard int = int(startCard)*20 + 20
    if endCard > len(qsls) {
	endCard = len(qsls)
    }
    var numInRow int = 0
    fmt.Fprintf(w, "<table boarder=\"1\" align=\"center\"><tr>")
    for i := int(startCard)*20; i < endCard; i++ {
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
    fmt.Fprintf(w, "<tr><td><div style=\"text-align:left\"><a href=/" + strconv.Itoa(newStart/20) + ">Back</a></div></td><td></td><td></td><td>  <div style=\"text-align:right\"> <a href=/" + strconv.Itoa(endCard/20) + ">Next</a></td></tr></table>")
}

func displayCard(w http.ResponseWriter, r *http.Request) {
    var callToCheck string = strings.ToUpper(r.URL.Path[6:])
    fmt.Fprintf(w, "<title>" + callToCheck + "</title>")
    fmt.Fprintf(w, "<div style=\"text-align:center\"><a href=/><img src=\"" + imagesFolder + logoFileName + "\" width=320 height=80></img></a>")
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
