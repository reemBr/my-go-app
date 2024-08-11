// main.go
package main

import (
    "net/http"
    "log"
)

func main() {
    // Serve static files from the current directory
    http.Handle("/", http.FileServer(http.Dir(".")))

    // Define the port on which the server will listen
    port := "8080"
    log.Printf("Starting server on:%s\n", port)
    
    // Start the server
    err := http.ListenAndServe(":"+port, nil)
    if err != nil {
        log.Fatal(err)
    }
}
