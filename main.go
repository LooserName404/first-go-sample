package main

import(
    "log"
    "net/http"
    "encoding/json"
)

type Message struct {
    Text string `json:text`
}

func homePage(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(Message{"Hello from the Home Page!"})
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":80", nil))
}

func main() {
    handleRequests()
}
