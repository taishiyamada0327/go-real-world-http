package main

import (
	"io"
	"log"
	"net/http"
)

// Http Get query
// func main() {
// 	values := url.Values{
// 		"query": {"hello world"},
// 	}

// 	resp, _ := http.Get("http://localhost:18888" + "?" + values.Encode())
// 	defer resp.Body.Close()
// 	body, _ := io.ReadAll(resp.Body)
// 	log.Println(string(body))
// 	log.Println("Fields :", resp.Header)
// }

// Http HEAD
func main() {
	resp, err := http.Head("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	body, _ := io.ReadAll(resp.Body)
	log.Println((string(body)))
	log.Println("Status :", resp.Status)
	log.Println("Headers : ", resp.Header)
}