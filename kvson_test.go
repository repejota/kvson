// Copyright 2016 The kvson Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kvson_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

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
