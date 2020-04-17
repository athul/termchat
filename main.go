package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

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
		StartServer()
	} else {
		runner()
	}

}
