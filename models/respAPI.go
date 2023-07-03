package models

import "github.com/aws/aws-lambda-go/events"

type RespApi struct {
	Status     int    `json:"status"`
	Message    string `json:"message"`
	CustomResp *events.APIGatewayProxyResponse
}
