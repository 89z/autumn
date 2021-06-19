package main

import "strings"

func scan(s string, sep byte) (string, string) {
	if n := strings.IndexByte(s, sep); n >= 0 {
		return s[:n], s[n+1:]
	}
	return s, ""
}

func main() {
	text, s := "", "north,east,south,west"
	for {
		text, s = scan(s, ',')
		if text+s == "" {
			break
		}
		println(text)
	}
}
