package main

import "strings"

func comma(r *strings.Reader) (string, error) {
	var s string
	for {
		b, e := r.ReadByte()
		if b == ',' {
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
		s, e := comma(r)
		if e != nil {
			break
		}
		println(s)
	}
}
