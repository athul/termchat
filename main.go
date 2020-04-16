package main

import (
	"flag"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	s "github.com/athul/termsocket/server"
)

var addr = flag.String("addr", ":8080", "http service address")

var prompt = &survey.Select{
	Message: "Choose an option:",
	Options: []string{"Server", "Messaging"},
	Default: "Server",
}

func main() {
	var replResponse string
	err := survey.AskOne(prompt, &replResponse)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if replResponse == "Server" {
		s.StartServer()
	} else if replResponse == "Messaging" {

	}

}
