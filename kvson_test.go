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

package kvson_test

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"

	. "github.com/repejota/kvson"
)

type Example struct {
	ID         int
	Key        string
	AnotherKey string
}

func TestInstancePath(t *testing.T) {
	var tmp = os.TempDir()
	kvson, err := NewKVSON(tmp)
	if err != nil {
		t.Error(err)
	}
	if kvson.Path != tmp {
		t.Errorf("Path is supposed to be '%s',  but found %s", tmp, kvson.Path)
	}
}

func TestInstanceUnexistingPath(t *testing.T) {
	var tmp = os.TempDir()
	var path = filepath.Join(tmp, "notexists")
	_, err := NewKVSON(path)
	if err.Error() != fmt.Sprintf("stat %s: no such file or directory", path) {
		t.Error("It should fail as path doesn't exist, but no error found.")
	}
}

func TestInstanceIsNotDirectoryPath(t *testing.T) {
	var tmp = os.TempDir()
	var path = filepath.Join(tmp, "notadirectory")
	ioutil.WriteFile(path, []byte("data"), 0644)
	_, err := NewKVSON(path)
	if err.Error() != fmt.Sprintf("stat %s: must be a directory", path) {
		t.Error("It should fail as path is not a directory, but no error found.")
	}
}

func TestPutString(t *testing.T) {
	var tmp = os.TempDir()
	kvson, err := NewKVSON(tmp)
	if err != nil {
		t.Error(err)
	}
	err = kvson.Put("foo", "bar")
	if err != nil {
		t.Errorf("It should not fail, but got an error: %s", err)
	}
}

func TestGetString(t *testing.T) {
	var tmp = os.TempDir()
	kvson, err := NewKVSON(tmp)
	if err != nil {
		t.Error(err)
	}
	var data string
	err = kvson.Get("foo", &data)
	if err != nil {
		t.Errorf("It should not fail, but got an error: %s", err)
	}
	if data != "bar" {
		t.Errorf("Value should be equal to 'bar', but got: %s", data)
	}
}

func TestPutStruct(t *testing.T) {
	var tmp = os.TempDir()
	kvson, err := NewKVSON(tmp)
	if err != nil {
		t.Error(err)
	}
	example := Example{
		ID:         1,
		Key:        "key",
		AnotherKey: "This is another key",
	}
	err = kvson.Put("example", example)
	if err != nil {
		t.Errorf("It should not fail, but got an error: %s", err)
	}
}

func TestGetStruct(t *testing.T) {
	var tmp = os.TempDir()
	kvson, err := NewKVSON(tmp)
	if err != nil {
		t.Error(err)
	}
	var example Example
	err = kvson.Get("example", &example)
	if err != nil {
		t.Errorf("It should not fail, but got an error: %s", err)
	}
	if example.ID != 1 {
		t.Errorf("Payload ID should be 1, but got: %d", example.ID)
	}
	if example.Key != "key" {
		t.Errorf("Payload ID should be 'key', but got: %s", example.Key)
	}
}

func TestExists(t *testing.T) {
	var tmp = os.TempDir()
	kvson, err := NewKVSON(tmp)
	kvson.Delete("exists")
	if err != nil {
		t.Error(err)
	}
	if kvson.Exists("exists") == true {
		t.Errorf("It should not exist, but it does")
	}
	err = kvson.Put("exists", "bar")
	if err != nil {
		t.Errorf("It should not fail, but got an error: %s", err)
	}
	if kvson.Exists("exists") == false {
		t.Errorf("It should exist, but it doesn't")
	}

	err = kvson.Delete("exists")
	if err != nil {
		t.Errorf("It should delete, but it doesn't")
	}
}

func TestDelete(t *testing.T) {

	var tmp = os.TempDir()
	kvson, err := NewKVSON(tmp)
	err = kvson.Delete("not_exists")
	if err == nil {
		t.Errorf("It should give an error. The key doesn't exists")
	}

	err = kvson.Put("exists_key", "foo")
	if err != nil {
		t.Errorf("It should not give an error. The key should be created")
	}

	err = kvson.Delete("exists_key")
	if err != nil {
		t.Errorf("It should not give an error. The key should be deleted")
	}

}

func BenchmarkPut(b *testing.B) {
	var tmp = os.TempDir()
	rand.Seed(time.Now().UnixNano())
	kvson, err := NewKVSON(tmp)
	for n := 0; n < b.N; n++ {
		err = kvson.Put(string(rand.Int()), n)
		if err != nil {
			b.Errorf("It should not fail, but got an error: %s", err)
		}
	}

}
