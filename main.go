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
    "os"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "path"
)

type qslObject struct {
    Callsign   string
    Date   string
    Front_image   string
    Back_image   string
    Mode   string
    Frequency string
}
var (
    qsls []qslObject
    rootdir, _ = os.Getwd()
)

func main() {
    file, e := ioutil.ReadFile(qsoFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }

    json.Unmarshal(file, &qsls)

    http.Handle("/compressedCards/", http.StripPrefix("/compressedCards",
        http.FileServer(http.Dir(path.Join(rootdir, convertedFolder)))))
    http.Handle("/thumbnails/", http.StripPrefix("/thumbnails",
	http.FileServer(http.Dir(path.Join(rootdir, resizedFolder)))))
    http.Handle("/cards/", http.StripPrefix("/cards",
        http.FileServer(http.Dir(path.Join(rootdir, cardsFolder)))))
    http.Handle("/images/", http.StripPrefix("/images",
        http.FileServer(http.Dir(path.Join(rootdir, imagesFolder)))))

    http.HandleFunc("/", index)
    http.HandleFunc("/view/", displayCard)
    http.HandleFunc("/api/", apiGetCall)


    fmt.Printf("Web Server started\n")
    http.ListenAndServe(":8080", nil)
}
