// Copyright 2016 The kvson Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package kvson implements a simple object notation storage module.
//
package kvson

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"os"
	"path/filepath"
)

// KVSON ...
type KVSON struct {
	Path string
}

// perm are the file permissions used
const perm os.FileMode = 0644

// getBytes converts an arbitrary Golang interface to byte array
func getBytes(payload interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(payload)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Get gets a value by key
func (s *KVSON) Get(key string) (payload []byte, err error) {
	filename := filepath.Join(s.Path, key)
	payload, err = ioutil.ReadFile(filename)
	return payload, err
}

// Put puts a value
func (s *KVSON) Put(key string, payload interface{}) error {
	filename := filepath.Join(s.Path, key)
	bytes, err := getBytes(payload)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, bytes, perm)
	if err != nil {
		return err
	}
	return nil
}

// NewKVSON allocates and initializes a new KVSON.
//
func NewKVSON(path string) *KVSON {
	return &KVSON{Path: path}
}
