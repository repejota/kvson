// Copyright 2016 The kvson Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kvson

import "testing"

func TestInstancePath(t *testing.T) {
	kvson := NewKVSON("/tmp")
	if kvson.Path != "/tmp" {
		t.Error("Path is supposed to be '/tmp',  but found", kvson.Path)
	}
}

func TestSaveString(t *testing.T) {
	kvson := NewKVSON("/tmp")
	err := kvson.Put("foo", "bar")
	if err != nil {
		t.Error("It should not fail, but got an error", err)
	}
}

func TestGetString(t *testing.T) {
	kvson := NewKVSON("/tmp")
	bytes, err := kvson.Get("foo")
	payload := string(bytes)
	if err != nil {
		t.Error("It should not fail, but got an error", err)
	}
	if payload != "bar" {
		t.Error("It should be equal to 'bar', but got", payload)
	}
}
