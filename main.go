package main

import (
	"net/http"

	"github.com/olahol/melody"
)

func main() {
	// Create a new melody instance to manage WebSocket connections
	m := melody.New()

	// Serve the index.html file at the root URL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// Handle WebSocket connections at the /ws endpoint
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// WebSocket handshake happens here
		m.HandleRequest(w, r)
	})

	// Broadcast received WebSocket messages to all connected clients
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})

	// Start the HTTP server on localhost at port 5000
	http.ListenAndServe("localhost:5000", nil)
}
