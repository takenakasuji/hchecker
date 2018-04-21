package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"github.com/fatih/color"
)

// Response struct
type Response struct {
	Status     string              `json:"status"`
	StatusCode int                 `json:"status_code"`
	Header     map[string][]string `json:"header"`
}

func main() {

	var (
		strURL = flag.String("url", "", "help message for \"url\" option")
		strFormat = flag.String("format", "inline", "help message for \"format\" option")
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

	res := Response{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
	}

	if *strFormat == "inline" {
		c := color.New(color.FgGreen, color.Bold)
		fmt.Printf("Status: ")
		c.Printf("%s\n", res.Status)
		fmt.Printf("StatusCode: ")
		c.Printf("%v\n", res.StatusCode)
		
		for key, value := range res.Header {
			fmt.Printf("%s: ", key)
			c.Printf("%s\n", value)
		}
	}

	if *strFormat == "json" {
		JSON, _ := json.Marshal(res)
		fmt.Println(string(JSON))
		return
	}
}
