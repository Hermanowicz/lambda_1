package main

import (
	"context"
	"encoding/json"

	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

type Response struct {
	Res string
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	ret := fmt.Sprintf("Hello, %s!", name.Name)
	r, _ := json.Marshal(Response{ret})
	return string(r), nil
}

func main() {
	lambda.Start(HandleRequest)
}
