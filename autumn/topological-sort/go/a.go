package main
import "fmt"

func main() {
   tertiary := map[string][]string{
      "citron": {"green", "orange"},
      "green": {"blue", "yellow"},
      "orange": {"red", "yellow"},
      "purple": {"red", "blue"},
      "russet": {"orange", "purple"},
      "slate": {"green", "purple"},
   }
   s := fmt.Sprint(tsort(tertiary, "slate"))
   if s == "[blue red purple yellow green slate]" {
      println(true)
   } else if s == "[blue yellow green red purple slate]" {
      println(true)
   } else {
      println(s)
   }
}

func tsort(graph map[string][]string, end string) []string {
   var (
      b = make(map[string]bool)
      l []string
      s = []string{end}
   )
   for len(s) > 0 {
      n := s[len(s) - 1]
      b[n] = true
      for _, m := range graph[n] {
         if ! b[m] {
            s = append(s, m)
         }
      }
      if s[len(s) - 1] == n {
         s = s[:len(s) - 1]
         l = append(l, n)
      }
   }
   return l
}
