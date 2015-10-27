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
