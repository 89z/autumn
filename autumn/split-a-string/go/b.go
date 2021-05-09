package main
import "strings"

type scanner struct {
   sep byte
   done bool
   text, todo string
}

func newScanner(todo string, sep byte) scanner {
   return scanner{todo: todo, sep: sep}
}

func (s *scanner) scan() bool {
   if s.done { return false }
   n := strings.IndexByte(s.todo, s.sep)
   if n == -1 {
      s.text, s.done = s.todo, true
   } else {
      s.text, s.todo = s.todo[:n], s.todo[n + 1:]
   }
   return true
}

func main() {
   s := newScanner("north,east,south,west", ',')
   for s.scan() {
      println(s.text)
   }
}
