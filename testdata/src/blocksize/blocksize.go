package blocksize

func longAsBlockSize() int {
	bs := 2
	return bs
}

func longerThanBlockSize() int {
	bs := 2
	bs++
	return bs // want "no blank line before"
}

func longerThanBlockSizeButWithEmptyLine() int {
	bs := 2
	return bs
}
