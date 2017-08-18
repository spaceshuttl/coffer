package router

import (
	"service/store"
	"testing"
)

type MockStore struct{}

func (ms *MockStore) All() ([]*store.Entry, error) {
	return []*store.Entry{
		&store.Entry{
			ID:    []byte("the id"),
			Key:   []byte("the key"),
			Value: []byte("the value"),
		},
	}, nil
}
func (ms *MockStore) Put(e *store.Entry) error    { return nil }
func (ms *MockStore) Delete(e *store.Entry) error { return nil }
func (ms *MockStore) AddCrypter(c *store.Crypter) { return }
func init() {
	dataStore = &MockStore{}
}

func TestHandleAll(t *testing.T) {
	msg := &Message{}
	resp, err := handleAll(msg)
	if err != nil {
		t.Error(err)
	}

	if len(resp) != 1 {
		t.Errorf("exptect respose length of %v got %v", 1, len(resp))
	}

	for _, entry := range resp {
		if recv := string(entry.ID); recv != "the id" {
			t.Errorf("expected %v got %v", "the id", recv)
		}
		if recv := string(entry.Key); recv != "the key" {
			t.Errorf("expected %v got %v", "the key", recv)
		}
		if recv := string(entry.Value); recv != "the value" {
			t.Errorf("expected %v got %v", "the value", recv)
		}
	}
}

func TestHandleAdd(t *testing.T) {
	msg := &Message{}
	resp, err := handleAdd(msg)
	if err != nil {
		t.Error(err)
	}

	if len(resp) != 1 {
		t.Errorf("exptect respose length of %v got %v", 1, len(resp))
	}

	for _, entry := range resp {
		if recv := string(entry.ID); recv != "the id" {
			t.Errorf("expected %v got %v", "the id", recv)
		}
		if recv := string(entry.Key); recv != "the key" {
			t.Errorf("expected %v got %v", "the key", recv)
		}
		if recv := string(entry.Value); recv != "the value" {
			t.Errorf("expected %v got %v", "the value", recv)
		}
	}
}

func TestHandleDelete(t *testing.T) {
	msg := &Message{}
	resp, err := handleDelete(msg)
	if err != nil {
		t.Error(err)
	}

	if len(resp) != 1 {
		t.Errorf("exptect respose length of %v got %v", 1, len(resp))
	}

	for _, entry := range resp {
		if recv := string(entry.ID); recv != "the id" {
			t.Errorf("expected %v got %v", "the id", recv)
		}
		if recv := string(entry.Key); recv != "the key" {
			t.Errorf("expected %v got %v", "the key", recv)
		}
		if recv := string(entry.Value); recv != "the value" {
			t.Errorf("expected %v got %v", "the value", recv)
		}
	}
}
