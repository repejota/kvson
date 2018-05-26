// Copyright 2018 Raül Pérez, repejota@gmail.com. All rights reserved.
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with this
// work for additional information regarding copyright ownership.  The ASF
// licenses this file to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package kvson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Version ...
var Version = "No version provided"

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

// Exists checks if a key is available on the store.
//
func (s *KVSON) Exists(key string) bool {
	filename := filepath.Join(s.Path, key)
	if _, err := os.Stat(filename); err == nil {
		return true
	}
	return false
}

// Delete if a exists a key on the store.
//
func (s *KVSON) Delete(key string) error {

	filename := filepath.Join(s.Path, key)
	_, err := os.Stat(filename)
	if err == nil {
		err = os.Remove(filename)
		if err != nil {
			return err
		}

	}

	return err
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
