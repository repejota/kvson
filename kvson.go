package kvson

import "io/ioutil"

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
func (e Element) Save(path string, id string, payload string) (err error) {
	filename := path + "/" + id
	err = ioutil.WriteFile(filename, []byte(payload), 0644)
	if err != nil {
		return err
	}
	return nil
}
