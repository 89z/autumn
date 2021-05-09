package main
import "strings"

type scanner struct { text, todo, sep string }

func newScanner(todo, sep string) scanner {
   return scanner{todo: todo, sep: sep}
}

func (s *scanner) scan() bool {
   if s.todo == "" { return false }
   a := strings.SplitN(s.todo, s.sep, 2)
   s.text = a[0]
   switch len(a) {
      case 1: s.todo = ""
      case 2: s.todo = a[1]
   }
   return true
}

func main() {
   s := newScanner("north,east,south,west", ",")
   for s.scan() {
      println(s.text)
   }
}
