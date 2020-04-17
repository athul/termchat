package main

import (
	"flag"

	"github.com/AlecAivazis/survey/v2"
)

var addr = flag.String("addr", ":8080", "http service address")

var prompt = &survey.Select{
	Message: "Choose an option:",
	Options: []string{"Server", "Messaging"},
	Default: "Server",
}

// func main() {
// 	// var replResponse string
// 	// err := survey.AskOne(prompt, &replResponse)
// 	// if err != nil {
// 	// 	fmt.Println(err.Error())
// 	// 	return
// 	// }
// 	// if replResponse == "Server" {
// 	// 	startServer()
// 	// }
// 	//StartServerplease()
// }
