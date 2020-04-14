package main

import (
	"flag"
	"log"

	s "github.com/athul/termsocket/server"
)

var addr = flag.String("addr", ":8080", "http service address")

// var prompt = &survey.Select{
// 	Message: "Choose an option:",
// 	Options: []string{"Server", "Messaging"},
// 	Default: "Server",
// }

func main() {
	// var replResponse string
	// err := survey.AskOne(prompt, &replResponse)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// if replResponse == "Server" {
	// 	startServer()
	// }
	log.Println("Server Started")
	s.StartServer()
}

// func startServer() {
// 	server := s.NewServer()
// 	go server.Listen()
// 	http.HandleFunc("/ws", handlews)

// 	http.ListenAndServe("localhost:8080", nil)
// }
// func handlews(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "chat.html")
// }
// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Home Page")
// }

// func main() {
// 	flag.Parse()
// 	hub := s.NewHub()
// 	go hub.Run()
// 	http.HandleFunc("/", homePage)
// 	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
// 		s.ServeWs(hub, w, r)
// 	})
// 	err := http.ListenAndServe(*addr, nil)
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }
