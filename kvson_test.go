package kvson

import "testing"

func TestGetError(t *testing.T) {
	el := Element{}
	_, err := el.Get("/tmp/foo")
	if err == nil {
		t.Error(err)
	}
}

func TestSave(t *testing.T) {
	el := Element{
		ID:      "test_id",
		Payload: []byte("test_payload"),
	}
	err := el.Save("/tmp")
	if err != nil {
		t.Error(err)
	}
}

func TestGet(t *testing.T) {
	el := Element{}
	_, err := el.Get("/tmp/test_id")
	if err != nil {
		t.Error(err)
	}
}
