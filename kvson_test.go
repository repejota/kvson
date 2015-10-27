package kvson

import "testing"

func TestDummy(t *testing.T) {

}

func TestSave(t *testing.T) {
	el := Element{
		ID:      "test_id",
		Payload: "test_payload",
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
