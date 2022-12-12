package padding

import (
	"math/rand"
	"testing"
)

func TestPad(t *testing.T) {
	input := []byte{45, 34, 22, 11, 22}
	padded, err := Pad(input, 1024)

	if err != nil {
		t.Error("Pad: ", err)
	}

	if len(padded) != 1024 {
		t.Errorf("len(Pad(input, 1024)) = %d, want 1024", len(padded))
	}
}

func TestUnpad(t *testing.T) {
	padded := make([]byte, 1022)
	n, err := rand.Read(padded)

	if err != nil {
		t.Error("rand.Read: ", err)
	}

	if n != len(padded) {
		t.Errorf("n = %d, want %d", n, len(padded))
	}

	padded = append(padded, 0x80)
	padded = append(padded, 0x00)
	unpadded, err := Unpad(padded, 1024)

	if err != nil {
		t.Error("Unpad: ", err)
	}

	if len(unpadded) != 1022 {
		t.Errorf("len(Pad(input, 1024)) = %d, want %d", len(unpadded), 1022)
	}
}

func TestUnpad_InvalidBlockSize(t *testing.T) {
	padded := []byte{45, 34, 22, 11, 22}
	_, err := Unpad(padded, 1024)

	if err == nil {
		t.Error("Unpad: must throw 'Invalid Padding' error when padded does not equal the block size")
	}
}

func TestUnpad_InvalidPadding(t *testing.T) {
	padded := make([]byte, 1022)
	n, err := rand.Read(padded)

	if err != nil {
		t.Error("rand.Read: ", err)
	}

	if n != len(padded) {
		t.Errorf("n = %d, want %d", n, len(padded))
	}

	padded = append(padded, 0x80)
	padded = append(padded, 0x01)
	_, err = Unpad(padded, 1024)

	if err == nil {
		t.Error("Unpad: must throw 'Invalid Padding' error when values other than 0x00 come after 0x80 in the padded data")
	}
}

func TestUnpad_InvalidPadding_NoMandatoryByte(t *testing.T) {
	padded := make([]byte, 1024)
	for i := range padded {
		padded[i] = 0x05
	}

	_, err := Unpad(padded, 1024)

	if err == nil {
		t.Error("Unpad: must throw 'Invalid Padding' error when no 0x80 is part of the padded data")
	}
}
