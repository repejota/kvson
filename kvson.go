package kvson

import (
	"fmt"
	"io/ioutil"
)

// Element ...
type Element struct {
	ID      string
	Payload string
}

// Get an element by its ID
func (e Element) Get(path string, id string) (el Element, err error) {
	filename := path + "/" + id
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return el, err
	}
	el = Element{
		ID:      id,
		Payload: string(data),
	}
	return el, nil
}

// Save an element
func (e Element) Save(path string) (err error) {
	filename := path + "/" + e.ID
	fmt.Println(filename)
	err = ioutil.WriteFile(filename, []byte(e.Payload), 0644)
	if err != nil {
		return err
	}
	return nil
}
