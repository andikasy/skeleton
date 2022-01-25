package skeleton

import "testing"

func TestCreate(test *testing.T) {
	directory := New()
	err := directory.Create()

	if err != nil {
		test.Error(test, err)
	}
}
