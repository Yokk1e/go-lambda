package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(reuqest events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Hello world")

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
	}

	return response, nil

}
