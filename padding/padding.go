package padding

import "errors"

var errNotMultipleBlockSize error = errors.New("length of padded input is not a multiple of block size")
var errinvalidPadding error = errors.New("invalid padding")

// Pad given input of length n adds a mandatory byte valued 0x80 to input followed by m < block-size bytes valued 0x00 such that n + m + 1 is the smallest multiple block-size (https://eris.codeberg.page/spec/#name-pad).
func Pad(input []byte, blockSize int) ([]byte, error) {
	n := len(input)

	// the final modulo blockSize ensures that if (n + 1) modulo block-size = 0 no unnecessary zeroes are added.
	m := (blockSize - ((n + 1) % blockSize)) % blockSize

	// add mandatory 0x80 followed by m bytes of zeroes (0x00)
	padded := append(append(input, 0x80), make([]byte, m)...)

	if (len(padded) % blockSize) != 0 {
		return []byte{}, errNotMultipleBlockSize
	}

	return padded, nil
}

// Unpad starts reading bytes from the end of input until a 0x80 is read and then returns bytes of input before the 0x80.
// This function throws an error if a value other than 0x00 is read before reading 0x80, if no 0x80 is read after reading block-size bytes from the end of input or if length of input is less than block-size (https://eris.codeberg.page/spec/#name-unpad).
func Unpad(input []byte, blockSize int) ([]byte, error) {

	n := len(input)

	// ensure that input is at least as large as block size
	if n < blockSize {
		return []byte{}, errinvalidPadding
	}

	for i := 1; i <= blockSize; i++ {

		// read the ith byte from the end of input
		byteFromInput := input[n-i]

		if byteFromInput == 0x80 {
			// special marker is reached, return everything before it
			return input[0:(n - i)], nil
		} else if byteFromInput == 0x00 {
			// continue with next byte from the right
			continue
		} else {
			// padding must be 0x00 or 0x80
			return []byte{}, errinvalidPadding
		}
	}

	// no 0x80 has been read after reading block-size bytes from the right of input
	return []byte{}, errinvalidPadding
}
