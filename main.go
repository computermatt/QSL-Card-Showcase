package main

import (
    "fmt"
    "os"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "path"
)

type jsonobject struct {
    Callsign   string
    Date   string
    Front_image   string
    Back_image   string
    Mode   string
    Frequency string
}

var (
    jsontype []jsonobject
    rootdir, _ = os.Getwd()
)

func main() {
    file, e := ioutil.ReadFile("../QSOs.json")
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }

    json.Unmarshal(file, &jsontype)

    apiGetCall()

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
