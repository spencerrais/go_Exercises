package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (Reader MyReader) Read(stream []byte) (int, error) {
	stream[0] = 'A'
	return 1, nil
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}
