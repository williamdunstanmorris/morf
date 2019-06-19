package main

import (
    "fmt"
    "net/http"
)

func main() {
       var port = 8000

       // this handler will recieve any incoming request
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
          // send a response text
          w.Write([]byte("Docker is awesome!!!"))
    })

        // print a message to the console
    fmt.Printf("Server running at http://127.0.0.1:%d", port)
        // start accepting new requests
        http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
