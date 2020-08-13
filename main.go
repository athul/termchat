package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/athul/termchat/msg"
)

var prompt = &survey.Select{
	Message: "Choose an option:",
	Options: []string{"Server", "Messaging"},
	Default: "Server",
}

func main() {

	setupCloseHandler()

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
func setupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		fmt.Println("Exiting from server.....")
		time.Sleep(1 * time.Second)
		os.Exit(0)
	}()
}
