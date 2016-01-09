// Copyright 2016 The kvson Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package kvson implements a simple object notation storage module.
//
package kvson

// KVSON ...
type KVSON struct {
	Path string
}

// element ...
type element struct {
	ID      string
	Payload interface{}
}

/*
// Get an element by its ID
func (s *KVSON) Get(key string) (el Element, err error) {
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
func (s *KVSON) Save(path string) (err error) {
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
*/

// NewKVSON allocates and initializes a new KVSON.
//
func NewKVSON(path string) *KVSON {
	kvson := KVSON{
		Path: path,
	}
	return &kvson
}
