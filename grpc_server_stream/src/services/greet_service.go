package services

import "fmt"

var GreetService greetServiceInterface = &greetService{}

type greetServiceInterface interface {
	Greet(firstName string, lastName string) string
}

type greetService struct{}

func (*greetService) Greet(firstName string, lastName string) string {
	return fmt.Sprintf("Hallo %s %s", firstName, lastName)
}
