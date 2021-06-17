package main

import "strings"

func scan(s, sep string) (string, string) {
	switch a := strings.SplitN(s, sep, 2); len(a) {
	case 0:
		return "", ""
	case 1:
		return a[0], ""
	}
	return a[0], a[1]
}

func main() {
	text, s := "", "north,east,south,west"
	for {
		text, s = scan(s, ",")
		if text+s == "" {
			break
		}
		println(text)
	}
}
