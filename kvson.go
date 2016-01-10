// Copyright 2016 The kvson Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package kvson implements a simple key-value object notation storage module.
//
// Keys are strings, and values can be any Golang type.
// All values are serialized to JSON when they are saved to disk. The key will
// be the name of the file.
//
package kvson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// KVSON is a type that holds the instance of a single key-value storage.
//
// All values are saved to a folder, so you can have multiple storage instances
// on different paths. It is not recommended to use multiple instances with the
// same path.
//
// The key for every element will be the name of the file under the base
// folder.
//
// The value can be any Golang type and will be serialized to JSON before
// writing data to disk.
//
type KVSON struct {
	Path string
}

// perm are the file permissions used to write values on the storage.
const (
	perm os.FileMode = 0644
)

// Get gets a value by key.
//
// It reads a file named as the key and unserializes the content of the file.
//
func (s *KVSON) Get(key string, value interface{}) error {
	filename := filepath.Join(s.Path, key)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &value)
	return err
}

// Put saves a value to the storage.
//
// It serializes a value to JSON and writes data to disk on a file named named
// as the key provided.
//
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
// It checks if the base path provided exists and it is a directory.
//
func NewKVSON(path string) (*KVSON, error) {
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil, err
	}
	if stat.Mode().IsRegular() {
		return nil, fmt.Errorf("stat %s: must be a directory", path)
	}
	return &KVSON{Path: path}, nil
}
