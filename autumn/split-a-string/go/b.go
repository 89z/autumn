package main

import (
	"bufio"
	"bytes"
)

func comma(b []byte, eof bool) (int, []byte, error) {
	if eof {
		return 0, nil, nil
	}
	if n := bytes.IndexByte(b, ','); n >= 0 {
		return n + 1, b[:n], nil
	}
	return len(b), b, nil
}

func main() {
	s := bufio.NewScanner(bytes.NewBufferString("north,east,south,west"))
	s.Split(comma)
	for s.Scan() {
		println(s.Text())
	}
}
