package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/athul/termchat/msg"
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
		msg.StartServer()
	} else if replResponse == "Messaging" {
		msg.Runner()
	} else {
		fmt.Print("Invalid Option")
	}

}
