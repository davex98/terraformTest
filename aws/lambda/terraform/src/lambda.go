package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
)

type AddRequest struct {
	A int  `json:"A"`
	B int `json:"B"`
}

func HandleRequest(ctx context.Context, req AddRequest) (int, error) {
	return req.A + req.B, nil
}

func main() {
	lambda.Start(HandleRequest)
}