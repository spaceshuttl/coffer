package router

import (
	"service/store"

	"github.com/sirupsen/logrus"
)

// TODO: implement
func handleLogin(m *Message) error {
	// logrus.WithFields(logrus.Fields{
	// 	"action":   m.Action,
	// 	"password": m.Payload.Password,
	// }).Debug("received websocket request")
	//
	// // Initialise our crypter
	// crypter, err := store.InitaliaseCrypter(m.Payload.Password)
	// if err != nil {
	// 	logrus.Error(err)
	// 	return err
	// }
	//
	// dataStore.AddCrypter(crypter)
	return nil
}

func handleAll(m *Message) ([]*store.Entry, error) {
	logrus.WithFields(logrus.Fields{
		"action": m.Action,
	}).Debug("received websocket request")

	return dataStore.All()
}

func handleAdd(m *Message) ([]*store.Entry, error) {
	logrus.WithFields(logrus.Fields{
		"action": m.Action,
	}).Debug("received websocket request")

	if err := dataStore.Put(m.Payload); err != nil {
		return nil, err
	}

	// Resend the updated store
	return dataStore.All()
}

func handleDelete(m *Message) ([]*store.Entry, error) {
	logrus.WithFields(logrus.Fields{
		"action": m.Action,
	}).Debug("received websocket request")

	if err := dataStore.Delete(m.Payload); err != nil {
		return nil, err
	}

	// Resend the updated store
	return dataStore.All()
}
