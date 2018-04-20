package main

import (
	"encoding/json"
	"fmt"
	"log"
    "net/http"
    "flag"
)

// Response struct
type Response struct {
    Status string `json:"status"`
    StatusCode int `json:"status_code"`
    Header map[string][]string `json:"header"`
}

func main() {

    var (
        strURL = flag.String("url", "", "help message for \"url\" option")
    )

    flag.Parse()

    if *strURL == "" {
        fmt.Println("need set url")
        return
    }

	resp, err := http.Get(*strURL)
	if err != nil {
		log.Fatal(err)
    }

    res := Response {
        Status: resp.Status,
        StatusCode: resp.StatusCode,
        Header: resp.Header,
    }

    JSON, _ := json.Marshal(res)
    fmt.Println(string(JSON))
}