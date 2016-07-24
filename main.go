package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const scope string = "https://www.googleapis.com/auth/cloud-platform"

const url string = "https://speech.googleapis.com"
const path string = "/v1beta1/speech:syncrecognize"

// Alternative is object holds transcribed text
type Alternative struct {
	Transcript string
	Confidence float64
}

// Result contains Alternative entries
type Result struct {
	Alternatives []Alternative
}

// Response contains results
type Response struct {
	Results []Result `json:results`
}

func main() {
	credentials := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	data, err := ioutil.ReadFile(credentials)
	if err != nil {
		log.Fatal(err)
	}
	conf, err := google.JWTConfigFromJSON(data, scope)
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(oauth2.NoContext)

	// TODO
	buf, err := os.Open("./sync-request.json")
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Post(url+path, "application/json", buf)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Fatal(res.Header)
		return
	}

	decoder := json.NewDecoder(res.Body)
	var response Response
	err = decoder.Decode(&response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.Results[0].Alternatives[0].Transcript)
}
