package main

import (
    "fmt"
    "net/http"
    "os"
    "net"
    "strings"
)

func root(w http.ResponseWriter, req *http.Request) {
    hostname, err := os.Hostname()
    if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }

    info, _ := net.InterfaceAddrs()
        ipx := strings.Split(info[1].String(), "/")[0]

    output := "<b>Hostname: </b>" + hostname + "</br>" + 
              "<b>IP address: </b>" + ipx + "</br>" +
              "<b>Version:</b> v5"
    fmt.Fprintf(w, output)
}

func main() {

    http.HandleFunc("/", root)

    http.ListenAndServe(":8080", nil)
}
