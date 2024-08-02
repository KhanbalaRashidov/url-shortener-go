package data

import (
	"fmt"
	"log"
)

type Store interface {
	Add(shortenedURL, longURL string) error
	Remove(shortenedURL string) error
	Get(shortendURL string) (string, error)
}

type MemoryStore struct {
	items map[string]string
}

func (m *MemoryStore) Add(shortendURL, longURL string) error {
	if m.items[shortendURL] != "" {
		return fmt.Errorf("value already exists here")
	}

	m.items[shortendURL] = longURL
	log.Println(m.items)

	return nil
}

func (m *MemoryStore) Remove(shortenedURL string) error {
	if m.items[shortenedURL] == "" {
		return fmt.Errorf("value does not exist here")
	}
	delete(m.items, shortenedURL)

	return nil
}

func (m *MemoryStore) Get(shortenedURL string) (string, error) {
	longURL, ok := m.items[shortenedURL]
	if !ok {
		return "", fmt.Errorf("no mapped url available here")
	}
	return longURL, nil
}

func NewMemoryStore() MemoryStore {
	return MemoryStore{items: make(map[string]string)}
}
