package router

import (
	"service/store"

	"github.com/sirupsen/logrus"
)

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
