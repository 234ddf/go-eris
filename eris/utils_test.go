package eris

import (
	"math/rand"
	"testing"
)

func TestSplitContent(t *testing.T) {
	input := make([]byte, 1025)
	n, err := rand.Read(input)

	if err != nil {
		t.Error("rand.Read: ", err)
	}

	if n != len(input) {
		t.Errorf("n = %d, want %d", n, len(input))
	}

	leafNodes, err := splitContent(input, 1024)

	if err != nil {
		t.Error("splitContent: ", err)
	}

	if len(leafNodes) != 2 {
		t.Errorf("len(leafNodes) = %d, want 2", len(leafNodes))
	}
}
