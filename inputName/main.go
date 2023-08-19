package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(Handler)
}


func Handler(event InputEvent) (string, error){
	fmt.Println(event.FirstName)
	fmt.Println(event.LastName)

	return "It worked!", nil
}

type InputEvent struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}