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

const (
	// ADD is the action to add an entry to the store
	ADD = "ADD"
	// ALL is the action to get all entries from the store
	ALL = "ALL"
	// GET is the action to get entries to the store
	GET = "GET"
	// DELETE is the action to delete an entry from the store
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
