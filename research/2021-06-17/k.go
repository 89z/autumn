package main

import "bytes"

func scan(r *bytes.Buffer, sep byte) string {
	var s string
	for {
		b := r.Next(1)
		if len(b) < 1 {
			if s != "" {
				break
			}
			return ""
		}
		if b[0] == sep {
			break
		}
		s += string(b)
	}
	return s
}

func main() {
	r := bytes.NewBufferString("north,east,south,west")
	for {
		s := scan(r, ',')
		if s == "" {
			break
		}
		println(s)
	}
}
