package dynamictime

import (
    "fmt"
    "time"
    "net/http"
)

func init() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world! It's me, Daniel.<br>")
    fmt.Fprint(w, "Server Time:  ")
    fmt.Fprintf(w, time.Now().Format(time.RFC1123))
}



