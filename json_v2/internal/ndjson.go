package internal

import (
	"bufio"
	"io"
)

func ScanLinesNoAlloc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := indexByte(data, '\n'); i >= 0 {
		return i + 1, data[:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}

func indexByte(b []byte, c byte) int {
	for i := range b {
		if b[i] == c {
			return i
		}
	}
	return -1
}

func NewNDJSONScanner(r io.Reader) *bufio.Scanner {
	s := bufio.NewScanner(r)
	buf := make([]byte, 0, 1024*1024)
	s.Buffer(buf, 64*1024*1024) // até 64 MB/linha para segurança
	s.Split(ScanLinesNoAlloc)
	return s
}
