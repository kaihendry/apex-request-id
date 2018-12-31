package main

import (
	"context"

	"github.com/apex/log"
	jsonhandler "github.com/apex/log/handlers/json"
	r "github.com/kaihendry/apex-request-id"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	log.SetHandler(jsonhandler.Default)
	lambda.Start(handler)
}

// Can we show the request ID of the lambda in every structured log?
func handler(ctx context.Context, evt events.SNSEvent) (string, error) {
	h := r.New(ctx)
	h.Log.Infof("Got the handle in Apex lambda environment")
	err := h.HellofromApex()
	return "", err
}
