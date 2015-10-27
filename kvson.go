package kvson

import (
	"io/ioutil"
	"path/filepath"
)

// Element ...
type Element struct {
	ID      string
	Payload string
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
		Payload: string(data),
	}
	return el, nil
}

// Save an element
func (e Element) Save(path string) (err error) {
	filename := filepath.Join(path, e.ID)
	err = ioutil.WriteFile(filename, []byte(e.Payload), 0644)
	if err != nil {
		return err
	}
	return nil
}
