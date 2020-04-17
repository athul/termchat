package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/fatih/color"
	"github.com/gorilla/websocket"
)

var (
	url     string
	origin  string
	red     = color.New(color.FgRed).SprintFunc()
	magenta = color.New(color.FgMagenta).SprintFunc()
	green   = color.New(color.FgGreen).SprintFunc()
	yellow  = color.New(color.FgYellow).SprintFunc()
	cyan    = color.New(color.FgCyan).SprintFunc()
	wg      sync.WaitGroup
)

func init() {
	flag.StringVar(&url, "url", "ws://localhost:8080/ws", "WebSocket server address to connect to")
}

func inLoop(ws *websocket.Conn, errors chan<- error, in chan<- []byte) {
	var msg = make([]byte, 512)
	for {
		//var n int
		var err error

		_, msg, err = ws.ReadMessage()
		//fmt.Println(n)
		if err != nil {
			errors <- err
			continue
		}

		in <- msg
	}

}

func printErrors(errors <-chan error) {
	for err := range errors {
		if err == io.EOF {
			fmt.Printf("\r✝ %v - connection closed by remote\n", magenta(err))
			os.Exit(0)
		} else {
			fmt.Printf("\rerr %v\n> ", red(err))
		}
	}
}

func printReceivedMessages(in <-chan []byte) {
	for msg := range in {
		fmt.Printf("\r< %s\n> ", cyan(string(msg)))
	}
}

func outLoop(ws *websocket.Conn, out <-chan []byte, errors chan<- error) {
	for msg := range out {
		err := ws.WriteMessage(1, msg)
		if err != nil {
			errors <- err
		}
	}
}

func runner() {
	flag.Parse()
	var d websocket.Dialer
	header := make(http.Header)
	header.Add("Origin", "http://localhost/")

	ws, _, err := d.Dial(url, header)

	defer ws.Close()

	if err != nil {
		panic(err)
	}

	fmt.Printf("successfully connected to %s\n\n", green(url))

	wg.Add(3)

	errors := make(chan error)
	in := make(chan []byte)
	out := make(chan []byte)

	defer close(errors)
	defer close(out)
	defer close(in)

	go inLoop(ws, errors, in)
	go printReceivedMessages(in)
	go printErrors(errors)
	go outLoop(ws, out, errors)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")
	for scanner.Scan() {
		out <- []byte(scanner.Text())

		fmt.Print("> ")
	}

	wg.Wait()
}
