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
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

type Index struct {
	Callsign   string
	ClubLogo   string
	Options    template.HTML
	RandomCard string
	TotalCards int
}

type Browse struct {
	Callsign string
	ClubLogo string
	QSLCards template.HTML
	BackCard string
	NextCard string
}

type Country struct {
	Callsign string
	ClubLogo string
	QSLCards template.HTML
}

type View struct {
	Callsign string
	ClubLogo string
	QSLCards template.HTML
}

func index(w http.ResponseWriter, r *http.Request) {

	opts := ""
	for _, element := range listOfContactedCountries {
		opts = opts + "<option>" + element + "</option>"
	}

	var randomCard int
	randomCard = int((rand.Float64()) * float64(len(qsls)))
	var fileName string = convertedFolder + qsls[randomCard].Front_image

	t, _ := template.ParseFiles("index.html")

	p := &Index{Callsign: callsign,
		ClubLogo:   imagesFolder + logoFileName,
		Options:    template.HTML(opts),
		RandomCard: fileName + convertedType,
		TotalCards: len(qsls)}

	t.Execute(w, p)

}

func browse(w http.ResponseWriter, r *http.Request) {
	var startCard, _ = strconv.ParseInt(r.URL.Path[8:], 0, 0)

	var endCard int = int(startCard)*20 + 20
	if endCard > len(qsls) {
		endCard = len(qsls)
	}
	var numInRow int = 0
	qslCards := ""
	for i := int(startCard) * 20; i < endCard; i++ {
		var call string = qsls[i].Callsign
		qslCards = qslCards + "<td><div style=\"text-align:center\"><a href=/view/" + call + ">" + call + "</a></br><img src=\"" + resizedFolder + qsls[i].Front_image + convertedType + "\" width=100 height=60></img>  <img src=\"" + resizedFolder + qsls[i].Back_image + convertedType + "\" width=100 height=60></img></div><p>     </p></td>"
		if numInRow == 3 {
			qslCards = qslCards + "</tr><tr>"
			numInRow = 0
		} else {
			numInRow++
		}
	}
	var newStart int = int(startCard) - 20
	if newStart <= 0 {
		newStart = 0
	}

	t, _ := template.ParseFiles("browse.html")

	p := &Browse{Callsign: callsign,
		ClubLogo: imagesFolder + logoFileName,
		QSLCards: template.HTML(qslCards),
		BackCard: strconv.Itoa(newStart / 20),
		NextCard: strconv.Itoa(endCard / 20)}

	t.Execute(w, p)

}

func browseCountry(w http.ResponseWriter, r *http.Request) {

	var cardsForCountry = listOfContactsPerCountry[strings.ToUpper(r.URL.Path[9:])]
	var numInRow int = 0
	qslCards := ""
	for i := 0; i < len(cardsForCountry); i++ {
		var call string = cardsForCountry[i].Callsign
		qslCards = qslCards + "<td><div style=\"text-align:center\"><a href=/view/" + call + ">" + call + "</a></br><img src=\"" + resizedFolder + cardsForCountry[i].Front_image + convertedType + "\" width=100 height=60></img>  <img src=\"" + resizedFolder + cardsForCountry[i].Back_image + convertedType + "\" width=100 height=60></img></div><p>     </p></td>"
		if numInRow == 3 {
			qslCards = qslCards + "</tr><tr>"
			numInRow = 0
		} else {
			numInRow++
		}
	}

	t, _ := template.ParseFiles("country.html")

	p := &Country{Callsign: callsign,
		ClubLogo: imagesFolder + logoFileName,
		QSLCards: template.HTML(qslCards)}

	t.Execute(w, p)
}

func displayCard(w http.ResponseWriter, r *http.Request) {
	var callToCheck string = strings.ToUpper(r.URL.Path[6:])
	qslCards := ""
	for i := 0; i < len(qsls); i++ {
		if qsls[i].Callsign == callToCheck {
			qslCards = qslCards + "<u><h1>" + qsls[i].Callsign + "</h1></u>"
			qslCards = qslCards + "<b>Country: " + PrefixLookup.CountryForCallsign(qsls[i].Callsign)
			qslCards = qslCards + "</br>Date: " + qsls[i].Date
			qslCards = qslCards + "</br>Mode: " + qsls[i].Mode
			qslCards = qslCards + "</br>Frequency: " + qsls[i].Frequency + "</b>"
			var fileName string = convertedFolder + qsls[i].Front_image
			qslCards = qslCards + "</br><table boarder=\"1\" align=\"center\"><tr><td><img src=\"" + fileName + convertedType + "\" width=480 height=320 ></img>"
			qslCards = qslCards + "</br><div style=\"text-align:center\"><a href=" + cardsFolder + qsls[i].Front_image + fullType + "> Download full sized image </a></div></td><td>"
			var backName string = "../" + convertedFolder + qsls[i].Back_image
			qslCards = qslCards + "<img src=\"" + backName + convertedType + "\" width=480 height=320></img>"
			qslCards = qslCards + "</br><div style=\"text-align:center\"><a href=../" + cardsFolder + qsls[i].Back_image + fullType + "> Download full sized image </a></div></td></tr>"
			qslCards = qslCards + "</table></br>"
		}
	}

	t, _ := template.ParseFiles("view.html")

	p := &View{Callsign: callsign,
		ClubLogo: imagesFolder + logoFileName,
		QSLCards: template.HTML(qslCards)}

	t.Execute(w, p)
}
