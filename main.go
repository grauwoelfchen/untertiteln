package main

import (
	//"encoding/json"
	//"fmt"
	//"io/ioutil"
	//"log"
	//"os"

	//"golang.org/x/oauth2"
	//"golang.org/x/oauth2/google"

	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"

	pb "github.com/grauwoelfchen/untertiteln/proto"
)

const scope string = "https://www.googleapis.com/auth/cloud-platform"

// rest
const url string = "https://speech.googleapis.com"
const path string = "/v1beta1/speech:syncrecognize"

// grpc
const address string = "speech.googleapis.com:443"
const method string = "/google.cloud.speech.v1beta1.Speech/SyncRecognize"

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
	keyFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	// REST syncrecognize
	//keyData, err := ioutil.ReadFile(keyFile)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//conf, err := google.JWTConfigFromJSON(keyData, scope)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//client := conf.Client(oauth2.NoContext)

	//buf, err := os.Open("./sync-request.json")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//res, err := client.Post(url+path, "application/json", buf)
	//defer res.Body.Close()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if res.StatusCode != 200 {
	//	log.Fatal(res.Header)
	//	return
	//}

	//decoder := json.NewDecoder(res.Body)
	//var response Response
	//err = decoder.Decode(&response)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(response.Results[0].Alternatives[0].Transcript)

	// gRPC SyncRecognize
	conf, err := oauth.NewServiceAccountFromFile(keyFile, scope)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
		grpc.WithPerRPCCredentials(conf),
		grpc.WithTimeout(5*time.Second),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	req := &pb.SyncRecognizeRequest{
		Config: &pb.RecognitionConfig{
			Encoding:     pb.RecognitionConfig_FLAC,
			SampleRate:   int32(16000),
			LanguageCode: "en-US",
		},
		Audio: &pb.RecognitionAudio{
			AudioSource: &pb.RecognitionAudio_Uri{
				Uri: "gs://cloud-samples-tests/speech/brooklyn.flac",
			},
		},
	}
	res := new(pb.SyncRecognizeResponse)
	ctx, cancel := context.WithTimeout(context.Background(), 12*time.Second)
	defer cancel()
	err = grpc.Invoke(ctx, method, req, res, conn)
	if err != nil {
		log.Fatal(err)
	}
	results := res.Results
	if len(results) > 0 {
		for i := 0; i < len(results); i++ {
			result := results[i]
			fmt.Println(result.Alternatives[0].Transcript)
		}
	}
}
