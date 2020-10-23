package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to the Home Page!</h1>")
}
func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Web Socket Endpoint :)")
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Successfully Connected...")
	reader(ws)
}
func HandleRequests() {
	//SetupRoutes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ws", wsEndpoint)
}
func main() {
	fmt.Println("Go WebSocket!")
	HandleRequests()
	// fmt.Println("Server is live on Port 3000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
