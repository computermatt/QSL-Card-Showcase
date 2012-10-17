package main

import (
    "fmt"
    "net/http"
    "path"
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


func startWebServer() {

    http.Handle("/convertedCards/", http.StripPrefix("/convertedCards",
        http.FileServer(http.Dir(path.Join(rootdir, "../convertedCards/")))))
    http.Handle("/cards/", http.StripPrefix("/cards",
	http.FileServer(http.Dir(path.Join(rootdir, "../cards/")))))
    http.Handle("/images/", http.StripPrefix("/images",
	http.FileServer(http.Dir(path.Join(rootdir, "../images/")))))

    http.HandleFunc("/", index)
    http.HandleFunc("/view/", displayCard)
    fmt.Printf("http server started\n\n")

    http.ListenAndServe(":8080", nil)
}

