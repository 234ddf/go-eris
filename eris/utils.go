package eris

func splitContent(content []byte, blockSize int) ([][]byte, error) {
	// pad content
	padded, err := pad(content, blockSize)

	if err != nil {
		return [][]byte{}, err
	}

	// length of padded content is a multiple of block size
	if (len(padded) % blockSize) != 0 {
		return [][]byte{}, errNotMultipleBlockSize
	}

	leafNodes := make([][]byte, len(padded)/blockSize)
	// split padded content
	for i := 0; i < len(leafNodes); i++ {
		leafNodes[i] = padded[(i * blockSize):((i + 1) * blockSize)]
	}
	return leafNodes, nil
}
