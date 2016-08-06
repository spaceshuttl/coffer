package router

import (
	"errors"
	"log"
	"net/http"
	"os/user"
	"sync"

	"service/store"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// Websocket actions
const (
	LOGIN  = "LOGIN"
	ADD    = "ADD"
	ALL    = "ALL"
	GET    = "GET"
	DELETE = "DELETE"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  2048,
		WriteBufferSize: 2048,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	wg sync.WaitGroup

	dataStore store.Datastore

	usr, _  = user.Current()
	baseDir = usr.HomeDir + "/.coffer"
)

// Start initialises the routes and started a listener
func Start(port string, ds *store.Store) error {
	dataStore = ds

	logrus.Debugf("starting router on port %s", port)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)

	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()
	for {
		var (
			m *Message
		)

		if err := conn.ReadJSON(&m); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				logrus.Error(err)
			}
			break
		}

		switch m.Action {
		case LOGIN:
			err := handleLogin(m)
			if err != nil {
				logrus.Error(err)
			}

			conn.WriteJSON(Response{
				Error: err,
			})

		case ALL:
			resp, err := handleAll(m)
			if err != nil {
				logrus.Error(err)
			}

			conn.WriteJSON(Response{
				Error:   err,
				Message: resp,
			})
		case ADD:
			resp, err := handleAdd(m)
			if err != nil {
				logrus.Error(err)
			}

			conn.WriteJSON(Response{
				Error:   err,
				Message: resp,
			})
		case DELETE:
			resp, err := handleDelete(m)
			if err != nil {
				logrus.Error(err)
			}

			conn.WriteJSON(Response{
				Error:   err,
				Message: resp,
			})
		default:
			conn.WriteJSON(Response{
				Error: ErrInvalidAction,
			})
		}
	}
}

// List of error messages
var (
	ErrInvalidAction = errors.New("invalid action")
)
