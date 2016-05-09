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

// Response is the universal format for WS responses
type Response struct {
	Data  []*store.Field `json:"data"`
	Error error          `json:"error"`
}

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
	logrus.Debugf("Received request %v:\n%v", r.URL.RequestURI(), r.Body)
	defer r.Body.Close()
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Error(err)
		return
	}
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			return
		}

		msg := store.Field{}
		if err := json.Unmarshal(p, &msg); err != nil {
			logrus.Error(err)
		}

		switch msg.Tag {

		case "ADD":
			addHandler(conn, &msg)

		case "DEL":
			deleteHandler(conn, &msg)

		case "ALL":
			all, err := allHandler()
			if err != nil {
				logrus.Error(err)
				conn.WriteJSON(&Response{
					Error: err,
				})
			}
			conn.WriteJSON(&Response{
				Data:  all,
				Error: nil,
			})
		default:
			conn.WriteJSON(&Response{
				Error: ErrInvalidTag,
			})
		}
	}
}

func addHandler(conn *websocket.Conn, msg *store.Field) {
	logrus.Debugf("adding %v to the store", *msg)

	err := Store.Put(msg)
	if err != nil {
		conn.WriteJSON(&Response{
			Error: err,
		})
		return
	}

	all, err := Store.GetAll()
	if err != nil {
		conn.WriteJSON(&Response{
			Error: err,
		})
		return
	}

	conn.WriteJSON(&Response{
		Data:  all,
		Error: nil,
	})
}

func getHandler(msg *store.Field) error {
	fmt.Printf("getting %v from the store", msg.Identifier)
	Store.Get(msg.Identifier)
	return nil
}

func deleteHandler(conn *websocket.Conn, msg *store.Field) {
	logrus.Debugf("deleting %+v from the store", *msg)

	if err := Store.Delete(msg.Identifier); err != nil {
		conn.WriteJSON(&Response{
			Error: err,
		})
		return
	}

	all, err := Store.GetAll()
	if err != nil {
		conn.WriteJSON(&Response{
			Error: err,
		})
		return
	}

	conn.WriteJSON(&Response{
		Data:  all,
		Error: nil,
	})
}

func allHandler() ([]*store.Field, error) {
	return Store.GetAll()
}

var (
	// ErrInvalidTag is the error returned when we receive a command with an unrecognised tag
	ErrInvalidTag = errors.New("error invalid command tag provided")
	// ErrNoPort is the error returned when no port is specified
	ErrNoPort = errors.New("error no port provided")
)
