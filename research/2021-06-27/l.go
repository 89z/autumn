package main

func fields(s string, sep rune) (string, string) {
	from := -1
	for to, r := range s {
		if r == sep {
			if from >= 0 {
				return s[from:to], s[to:]
			}
		} else if from < 0 {
			from = to
		}
	}
	if from >= 0 {
		return s[from:], ""
	}
	return "", ""
}

func main() {
	text, s := "", ",north,,south,"
	for {
		text, s = fields(s, ',')
		if text+s == "" {
			break
		}
		println(text)
	}
}
