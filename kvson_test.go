// Copyright 2016 The kvson Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kvson

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
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
		t.Error("Path is supposed to be '"+tmp+"',  but found", kvson.Path)
	}
}

func TestInstanceUnexistingPath(t *testing.T) {
	var tmp = os.TempDir()
	var path = filepath.Join(tmp, "notexists")
	_, err := NewKVSON(path)
	if err.Error() != "stat "+path+": no such file or directory" {
		t.Error("It should fail as path doesn't exist, but we don't find any error.")
	}
}

func TestInstanceIsNotDirectoryPath(t *testing.T) {
	var tmp = os.TempDir()
	var path = filepath.Join(tmp, "notadirectory")
	ioutil.WriteFile(path, []byte("data"), 0644)
	_, err := NewKVSON(path)
	if err.Error() != "stat "+path+": must be a directory" {
		t.Error("It should fail as path doesn't exist, but we don't find any error.")
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
		t.Error("It should not fail, but got an error", err)
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
		t.Error("It should not fail, but got an error", err)
	}
	if data != "bar" {
		t.Error("Value should be equal to 'bar', but got", data)
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
		t.Error("It should not fail, but got an error", err)
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
		t.Error("It should not fail, but got an error", err)
	}
	if example.ID != 1 {
		t.Error("Payload ID should be 1, but got", example.ID)
	}
	if example.Key != "key" {
		t.Error("Payload ID should be 'key', but got", example.Key)
	}
}
