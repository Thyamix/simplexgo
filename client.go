package simplexgo

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

// TODO: The client need to be thought out more at some point, this is mostly for testing at this point

func InitCliSocket() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:5225", nil)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				return
			}
			var event Event
			err = json.Unmarshal(message, &event)
			if err != nil {
				log.Printf("Error: %w", err)
			}
		}
	}()

	for {
	}
}
