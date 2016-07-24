package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const scope string = "https://www.googleapis.com/auth/cloud-platform"
const url string = "https://speech.googleapis.com"
const path string = "/v1beta1/speech:syncrecognize"

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
	_, err = io.Copy(os.Stdout, res.Body)
	if err != nil {
		log.Fatal(err)
	}
}
