package main

import (
	"strings"
	"text/scanner"
)

func main() {
	var s scanner.Scanner
	s.Init(strings.NewReader("Wonder.Woman.1984.2020.IMAX"))
	s.Whitespace = 1 << '.'
	for s.Scan() != scanner.EOF {
		println(s.TokenText())
	}
}
