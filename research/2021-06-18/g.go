package main

import "strings"

func scan(r *strings.Reader, sep byte) (string, error) {
	var s string
	for {
		b, e := r.ReadByte()
		if b == sep {
			break
		}
		if e != nil {
			if s != "" {
				break
			}
			return "", e
		}
		s += string(b)
	}
	return s, nil
}

func main() {
	r := strings.NewReader("north,east,south,west")
	for {
		s, e := scan(r, ',')
		if e != nil {
			break
		}
		println(s)
	}
}
