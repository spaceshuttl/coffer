package router

import "service/store"

// Message is the structure of a message we will send and receive over websocket
type Message struct {
	Action  string       `json:"action"`
	Payload *store.Entry `json:"payload"`
}

// Response is the strict message structure in which we send responses to the client
type Response struct {
	Error   error          `json:"error"`
	Message []*store.Entry `json:"message"`
}
