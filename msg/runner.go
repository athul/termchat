package msg

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh/terminal"
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
	//ClientName is the Name of the Client that is exported
	colors = map[int]func(a ...interface{}) string{1: red, 2: magenta, 3: green, 4: yellow, 5: cyan}
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
			os.Exit(0)
		}
	}
}

func printReceivedMessages(in <-chan []byte) {
	w, _, _ := terminal.GetSize(0)

	for msg := range in {
		fmt.Printf("\r< %s"+toRight(">", w), string(msg))
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
func getname() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your Name:  ")
	ClientName, _ := reader.ReadString('\n')
	if ClientName == "" {
		fmt.Println("You need to have a name to join the Chat")
		ClientName, _ := reader.ReadString('\n')
		log.Println(ClientName)
		return ClientName
	}
	ClientName = strings.TrimSpace(ClientName)
	log.Println(ClientName)
	return ClientName

}

// Runner starts the messaging client
func Runner() {
	name := getname()
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
	//Logserver(fmt.Sprintf(`%s joined the Chat`, name))
	fmt.Printf("%s joined the Chat\n", name)
	fmt.Printf("[%s]>", name)
	namecolor := colors[selectRandom()]
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			// out <- []byte(fmt.Sprintf(`[%s] %s`, namecolor(name), "---nil---"))
			fmt.Printf(`        >`)
		} else {
			resp := fmt.Sprintf(`[%s] %s`, strings.TrimSpace(namecolor(name)), namecolor(text))
			out <- []byte(resp)
		}
	}
	wg.Wait()
}
func toRight(s string, w int) string {
	return fmt.Sprintf("%"+strconv.Itoa(w)+"s  ", s)
}
func selectRandom() int {
	rand.Seed(time.Now().Unix())
	ints := []int{1, 2, 3, 4, 5}
	randomIndex := rand.Intn(len(ints))
	pick := ints[randomIndex]
	// Logserver(pick)
	return pick
}

//Logserver Prints to the Server Rather than terminal
func Logserver(a ...interface{}) {
	log.Println(green(a...))
}
