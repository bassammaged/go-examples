package stack

import (
	"testing"
)

func TestAddValue(t *testing.T) {
	rs := RandomStack{}
	rs.Push("Kemet")

	if len(rs.items) != 0 {
		t.Log("Success")
	} else {
		t.Errorf("Failed")
	}
}
