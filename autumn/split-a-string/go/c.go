package main
import "strings"
type scanner struct { s, sep, text string }

func newScanner(s, sep string) scanner {
   return scanner{s: s, sep: sep}
}

func (s *scanner) scan() bool {
   if s.s == "" { return false }
   a := strings.SplitN(s.s, s.sep, 2)
   s.text, s.s = a[0], ""
   if len(a) > 1 {
      s.s = a[1]
   }
   return true
}

func main() {
   s := newScanner("north,east,south,west", ",")
   for s.scan() {
      println(s.text)
   }
}
