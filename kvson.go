// Copyright 2016 The kvson Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package kvson implements a simple object notation storage module.
//
package kvson

import (
	"encoding/json"
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

// Get gets a value by key
func (s *KVSON) Get(key string, value interface{}) error {
	filename := filepath.Join(s.Path, key)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &value)
	return err
}

// Put puts a value
func (s *KVSON) Put(key string, value interface{}) error {
	filename := filepath.Join(s.Path, key)
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, data, perm)
	if err != nil {
		return err
	}
	return nil
}

// NewKVSON allocates and initializes a new KVSON.
//
func NewKVSON(path string) (*KVSON, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}
	return &KVSON{Path: path}, nil
}
