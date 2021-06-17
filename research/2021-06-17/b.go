package main

import (
	"bufio"
	"strings"
)

func scan(r *bufio.Reader, sep byte) (string, error) {
	s, e := r.ReadString(sep)
	if s == "" {
		return "", e
	}
	if e != nil {
		return s, nil
	}
	return s[:len(s)-1], nil
}

func main() {
	r := bufio.NewReader(strings.NewReader("north,east,south,west"))
	for {
		s, e := scan(r, ',')
		if e != nil {
			break
		}
		println(s)
	}
}
