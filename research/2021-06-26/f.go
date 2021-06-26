package main

import (
	"strings"
	"text/scanner"
)

func main() {
	var (
		s scanner.Scanner
		w = "Wonder.Woman.1984.2020.IMAX.1080p.WEBRip.x265-RARBG"
	)
	s.Init(strings.NewReader(w))
	s.Mode = scanner.ScanIdents | scanner.ScanInts
	s.Whitespace = 1 << '.'
	for s.Scan() != scanner.EOF {
		println(s.TokenText())
	}
}
