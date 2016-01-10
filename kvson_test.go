// Copyright 2016 The kvson Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kvson

import "testing"

type Example struct {
	ID         int
	Key        string
	AnotherKey string
}

func TestInstancePath(t *testing.T) {
	kvson := NewKVSON("/tmp")
	if kvson.Path != "/tmp" {
		t.Error("Path is supposed to be '/tmp',  but found", kvson.Path)
	}
}

func TestPutString(t *testing.T) {
	kvson := NewKVSON("/tmp")
	err := kvson.Put("foo", "bar")
	if err != nil {
		t.Error("It should not fail, but got an error", err)
	}
}

func TestGetString(t *testing.T) {
	kvson := NewKVSON("/tmp")
	var data string
	err := kvson.Get("foo", &data)
	if err != nil {
		t.Error("It should not fail, but got an error", err)
	}
	if data != "bar" {
		t.Error("Value should be equal to 'bar', but got", data)
	}
}

func TestPutStruct(t *testing.T) {
	example := Example{
		ID:         1,
		Key:        "key",
		AnotherKey: "This is another key",
	}
	kvson := NewKVSON("/tmp")
	err := kvson.Put("example", example)
	if err != nil {
		t.Error("It should not fail, but got an error", err)
	}
}

func TestGetStruct(t *testing.T) {
	kvson := NewKVSON("/tmp")
	var example Example
	err := kvson.Get("example", &example)
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
