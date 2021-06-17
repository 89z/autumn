package main

import (
	"bufio"
	"strings"
)

func scan(b *bufio.Scanner) string {
	var s string
	for b.Scan() {
		if b.Text() == "," {
			if s == "" {
				continue
			}
			break
		}
		s += b.Text()
	}
	return s
}

func main() {
	b := bufio.NewScanner(strings.NewReader("north,east,south,west"))
	b.Split(bufio.ScanRunes)
	for {
		s := scan(b)
		if s == "" {
			break
		}
		println(s)
	}
}
