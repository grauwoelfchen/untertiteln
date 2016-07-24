#!/bin/sh

GCLOUD="${GOPATH}/google-cloud-sdk/bin/gcloud"
API_URL="https://speech.googleapis.com/v1beta1/speech:syncrecognize"
AUDIO_FILE="sync-request.json"

ACCOUNT=$($GCLOUD auth activate-service-account \
  --key-file=${GOOGLE_APPLICATION_CREDENTIALS})
echo $ACCOUNT

TOKEN=$($GCLOUD auth print-access-token)
echo $TOKEN

curl -s -k \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ${TOKEN}" \
  $API_URL \
  -d @${AUDIO_FILE}
