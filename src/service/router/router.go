package router

import (
	"log"
	"net/http"
	"sync"

	"service/store"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	wg sync.WaitGroup

	// ADD is the action to add an entry to the store
	ADD = "ADD"
	// ALL is the action to get all entries from the store
	ALL = "ALL"
	// GET is the action to get entries to the store
	GET = "GET"
	// DELETE is the action to delete an entry from the store
	DELETE = "DELETE"

	dataStore *store.Store
)

// Message is the structure of a message we will send and receive over websocket
type Message struct {
	Action  string       `json:"action"`
	Payload *store.Entry `json:"payload"`
}

// Start initialises the routes and started a listener
func Start(port string, str *store.Store) error {
	dataStore = str

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

	wg.Add(1)
	go connhandler(conn)
}

func connhandler(conn *websocket.Conn) {
	for {
		var m *Message
		err := conn.ReadJSON(&m)
		if err != nil {
			return
		}

		switch m.Action {
		case ALL:
			logrus.Debugf("Got WS action: %s", m.Action)
			entries, err := dataStore.All()
			if err != nil {
				logrus.Error(err)
				conn.WriteJSON(err)
			}
			logrus.Debugf("sending entries %+v", entries)
			conn.WriteJSON(entries)

		case ADD:
			logrus.Debugf("Got WS action: %s", m.Action)
			err := dataStore.Put(m.Payload)
			if err != nil {
				logrus.Error(err)
				conn.WriteJSON(err)
			}
			// HACK: resend them the updated store
			entries, err := dataStore.All()
			if err != nil {
				logrus.Error(err)
				conn.WriteJSON(err)
			}
			conn.WriteJSON(entries)

		case DELETE:
			logrus.Debugf("Got WS action: %s", m.Action)
			err := dataStore.Delete(m.Payload)
			if err != nil {
				logrus.Error(err)
				conn.WriteJSON(err)
			}
			// HACK: resend them the updated store
			entries, err := dataStore.All()
			if err != nil {
				logrus.Error(err)
				conn.WriteJSON(err)
			}
			conn.WriteJSON(entries)

			// case GET:
			// logrus.Debugf("Got WS action: %s", m.Action)
			// entries, err := dataStore.All()
			// if err != nil {
			// 	conn.WriteJSON(err)
			// }
			// conn.WriteJSON(entries)
		}

	}
}
