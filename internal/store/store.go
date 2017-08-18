package store

type Store interface {
	Put(*Entity) error
	Get(key string) (*Entity, error)
	Flush() error
}