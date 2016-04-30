package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"service/store"

	"github.com/gorilla/websocket"
)

var (
	// Store holds the local store
	Store    *store.Store
	upgrader = websocket.Upgrader{
		// HACK(mnzt): We want to check origin for security reasons.
		CheckOrigin: func(r *http.Request) bool { return true },
	}
)

// Init will initialise the API routes
func Init(port string, store *store.Store) error {

	if port == "" {
		return ErrNoPort
	}

	// Set our local store
	Store = store

	// Deal with API routes
	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(":5050", nil); err != nil {
		// HACK(mnzt): KEEP CALM AND DON'T PANIC
		panic(err)
	}

	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return
		}

		msg := store.Field{}

		if err := json.Unmarshal(p, &msg); err != nil {
			panic(err)
		}

		switch msg.Tag {
		case "add":
			err := addHandler(&msg)
			if err != nil {
				conn.WriteMessage(messageType, []byte("err"))
			}
		case "delete":
			err := deleteHandler(&msg)
			if err != nil {
				conn.WriteMessage(messageType, []byte("err"))
			}
		case "all":
			all, err := allHandler()
			if err != nil {
				conn.WriteMessage(messageType, []byte("err"))
			}
			ba, _ := json.Marshal(all)
			conn.WriteMessage(messageType, ba)
		default:
			conn.WriteMessage(messageType, []byte("invalid tag"))
		}
	}
}

func addHandler(msg *store.Field) error {
	// store password
	fmt.Printf("adding %+v to the store", msg)

	err := Store.Put(msg)
	if err != nil {
		return err
	}

	return nil
}

func getHandler(msg *store.Field) error {
	fmt.Printf("getting %v from the store", msg.Identifier)
	Store.Get(msg.Identifier)
	return nil
}

func deleteHandler(msg *store.Field) error {
	//delete passwor
	fmt.Printf("deleting %+v from the store", msg)
	return nil
}

func allHandler() ([]*store.Field, error) {
	return Store.GetAll()
}

var (
	// ErrNoPort is the error returned when no port is specified
	ErrNoPort = errors.New("error no port provided")
)
