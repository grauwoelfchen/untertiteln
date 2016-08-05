package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"

	pb "github.com/grauwoelfchen/untertiteln/proto/sync_recognize"
)

const (
  scope string = "https://www.googleapis.com/auth/cloud-platform"
  address string = "speech.googleapis.com:443"
  method string = "/google.cloud.speech.v1beta1.Speech/SyncRecognize"
)

func main() {
	keyFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

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
