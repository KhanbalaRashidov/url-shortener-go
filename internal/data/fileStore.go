package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type internalStore struct {
	Version string            `json:"version"`
	Items   map[string]string `json:"items"`
}

type FileStore struct {
	filename string
}

func NewFileStore(filename string) (FileStore, error) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		is := internalStore{Version: "v1", Items: make(map[string]string)}
		raw, err := json.Marshal(is)
		if err != nil {
			return FileStore{}, fmt.Errorf("unable to generate json representation for file")
		}
		err = ioutil.WriteFile(filename, raw, 0644)
		if err != nil {
			return FileStore{}, fmt.Errorf("unable to persist file")
		}
	}
	return FileStore{filename: filename}, nil
}

func (f *FileStore) Add(shortendURL, longURL string) error {
	raw, err := ioutil.ReadFile(f.filename)
	if err != nil {
		return err
	}

	var is internalStore
	err = json.Unmarshal(raw, &is)
	if err != nil {
		return fmt.Errorf("unable to parse incoming json store data. Err: %v", err)
	}

	_, ok := is.Items[shortendURL]
	if ok {
		return fmt.Errorf("shortened url already stored")
	}
	is.Items[shortendURL] = longURL
	modRaw, err := json.Marshal(is)
	if err != nil {
		return fmt.Errorf("unable to convert data to json representation")
	}

	err = ioutil.WriteFile(f.filename, modRaw, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (f *FileStore) Remove(shortenedURL string) error {
	raw, err := ioutil.ReadFile(f.filename)

	if err != nil {
		return err
	}
	var is internalStore
	err = json.Unmarshal(raw, &is)
	if err != nil {
		return fmt.Errorf("unable to parse incoming json store data. Err: %v", err)
	}
	delete(is.Items, shortenedURL)
	modRaw, err := json.Marshal(is)
	if err != nil {
		return fmt.Errorf("unable to convert data to json representation")
	}
	err = ioutil.WriteFile(f.filename, modRaw, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (f *FileStore) Get(shortendURL string) (string, error) {
	raw, err := ioutil.ReadFile(f.filename)

	if err != nil {
		return "", err
	}
	var is internalStore
	err = json.Unmarshal(raw, &is)
	if err != nil {
		return "", fmt.Errorf("unable to parse incoming json store data. Err: %v", err)
	}
	longURL, ok := is.Items[shortendURL]
	if !ok {
		return "", fmt.Errorf("no url available for that shortened url")
	}
	return longURL, nil
}
