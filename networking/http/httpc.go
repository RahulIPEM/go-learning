package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Job is a job description
type Job struct {
	User   string `json:"user"`
	Action string `json:"action"`
	Count  int    `json:"count"`
}

func main() {
	// GET request
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalf("error: can't call httpbin.org")
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)

	fmt.Println("-------------")

	job := &Job{
		User:   "Rahul",
		Action: "punch",
		Count:  1,
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(job); err != nil {
		log.Fatalf("error: can't encode job - %s", err)
	}

	resp1, err1 := http.Post("https://httpbin.org/post", "application/json", &buf)
	if err1 != nil {
		log.Fatalf("error: can't call httpbin.org")
	}
	defer resp1.Body.Close()

	io.Copy(os.Stdout, resp1.Body)
}
