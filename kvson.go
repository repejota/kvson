package kvson

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

// Element ...
type Element struct {
	ID      string
	Payload interface{}
}

// Get an element by its ID
func (e Element) Get(path string) (el Element, err error) {
	base := filepath.Base(path)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return el, err
	}
	el = Element{
		ID:      base,
		Payload: data,
	}
	return el, nil
}

// Save an element
func (e Element) Save(path string) (err error) {
	filename := filepath.Join(path, e.ID)
	payload, err := json.Marshal(e.Payload)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, payload, 0644)
	if err != nil {
		return err
	}
	return nil
}
