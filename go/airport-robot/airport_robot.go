package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greater interface {
	LanguageName() string
	Greet(name string) string
}

type Italian struct{}
type Portuguese struct{}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s", name)
}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func (p Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s", name)
}

func SayHello(name string, g Greater) string {
	return fmt.Sprintf("I can speak %s: %s!", g.LanguageName(), g.Greet(name))
}
