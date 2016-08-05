package router

import (
	"log"
	"net/http"
	"os/user"
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

	usr, _  = user.Current()
	baseDir = usr.HomeDir + "/coffer"
)

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

// Start initialises the routes and started a listener
func Start(port string, str *store.Store) error {
	dataStore = str

	http.HandleFunc("/", handler)

	// http.ListenAndServe(":"+port, nil)

	// TODO(mnzt): serve WS over TLS
	// certFile, keyFile, err := generateCerts()
	// if err != nil {
	// 	return err
	// }
	// http.ListenAndServeTLS(":"+port, certFile, keyFile, nil)
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
				conn.WriteJSON(Response{
					Error:   err,
					Message: entries,
				})
			}
			logrus.Debugf("sending entries %+v", entries)
			conn.WriteJSON(Response{
				Error:   err,
				Message: entries,
			})

		case ADD:
			logrus.Debugf("Got WS action: %s", m.Action)
			err := dataStore.Put(m.Payload)
			if err != nil {
				logrus.Error(err)
				conn.WriteJSON(Response{
					Error: err,
				})
			}
			// HACK: resend them the updated store
			entries, err := dataStore.All()
			if err != nil {
				logrus.Error(err)
				conn.WriteJSON(Response{
					Error:   err,
					Message: entries,
				})
			}
			conn.WriteJSON(Response{
				Error:   err,
				Message: entries,
			})
		case DELETE:
			logrus.Debugf("Got WS action: %s", m.Action)
			err := dataStore.Delete(m.Payload)
			if err != nil {
				logrus.Error(err)
				conn.WriteJSON(Response{
					Error: err,
				})
			}

			// HACK: resend them the updated store
			entries, err := dataStore.All()
			if err != nil {
				logrus.Error(err)
				conn.WriteJSON(Response{
					Error:   err,
					Message: entries,
				})
			}
			conn.WriteJSON(Response{
				Error:   err,
				Message: entries,
			})
		}
	}
}

// func generateCerts() (string, string, error) {
// 	var (
// 		certFile = baseDir + "/cert.pem"
// 		keyFile  = baseDir + "/key.pem"
// 	)
// 	// rand.Seed(time.Now().Unix())
// 	priv, err := rsa.GenerateKey(rand.Reader, 4096)
// 	if err != nil {
// 		return "", "", err
// 	}
// 	publickey := &priv.PublicKey
//
// 	c := &x509.Certificate{
// 		IsCA: true,
// 	}
//
// 	var parent = c
// 	// Create a self signed certificate
// 	cert, err := x509.CreateCertificate(rand.Reader, c, parent, publickey, priv)
// 	if err != nil {
// 		return "", "", err
// 	}
//
// 	// pKey := x509.MarshalPKCS1PrivateKey(cert)
// 	err = ioutil.WriteFile(certFile, cert, 0666)
// 	if err != nil {
// 		return "", "", err
// 	}
//
// 	pubKey, err := x509.MarshalPKIXPublicKey(publickey)
// 	if err != nil {
// 		return "", "", err
// 	}
//
// 	err = ioutil.WriteFile(keyFile, pubKey, 0666)
// 	if err != nil {
// 		return "", "", err
// 	}
//
// 	return certFile, keyFile, nil
// }
