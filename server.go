package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
)

func main() {
    fmt.Printf("Starting server at port 3000\n")

    http.HandleFunc("/", handleWebhook)
    http.HandleFunc("/hello", sayHello)

    if err := http.ListenAndServe(":3000", nil); err != nil {
        log.Fatal(err)
    }
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
    fmt.Println("hello there!")

    if r.URL.Path != "/" {
        http.Error(w, "404!", http.StatusNotFound)
        return
    }

    samplePayload := `{"modified": ["sample.json" ]}`

    var item map[string]interface{}

    json.Unmarshal([]byte(samplePayload), &item)
    fmt.Println(item)

    if r.Method == "POST" {
        fmt.Println("in the POST!");
    }
}

func sayHello(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("hello!!")

}