package main

import (
   "golang.org/x/text/language"
   "golang.org/x/text/message"
)

func main() {
   n := 1000
   s := message.NewPrinter(language.English).Sprint(n)
   println(s == "1,000")
}
