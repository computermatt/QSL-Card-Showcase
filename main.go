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
    file, e := ioutil.ReadFile(qsoFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }

    json.Unmarshal(file, &jsontype)

    http.Handle("/convertedCards/", http.StripPrefix("/convertedCards",
        http.FileServer(http.Dir(path.Join(rootdir, convertedFolder)))))
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
