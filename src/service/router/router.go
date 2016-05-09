package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"service/store"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var (
	// certFile = "./cert.pem"
	// keyFile  = "./key.pem"

	// Store holds the local store
	Store *store.Store

	upgrader = websocket.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 4096,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// Init will initialise the API routes
func Init(port string, store *store.Store) error {
	logrus.Info("Starting web server...")
	if port == "" {
		return ErrNoPort
	}

	// Set our local store
	Store = store

	// Deal with API routes
	http.HandleFunc("/", handler)

	// if err := http.ListenAndServeTLS(":"+port, certFile, keyFile, nil); err != nil {
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logrus.Infof("Starting server on %s", port)
		return err
	}

	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	logrus.Debugf("Received request %v", r.URL.RequestURI())
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
	logrus.Debugf("adding %v to the store", *msg)

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
	//delete password
	logrus.Debugf("deleting %+v from the store", *msg)
	Store.Delete(msg.Identifier)
	return nil
}

func allHandler() ([]*store.Field, error) {
	return Store.GetAll()
}

var (
	// ErrNoPort is the error returned when no port is specified
	ErrNoPort = errors.New("error no port provided")
)
