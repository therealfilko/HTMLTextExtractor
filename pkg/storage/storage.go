package storage

type Storage interface {
    Save(data string) error
}
