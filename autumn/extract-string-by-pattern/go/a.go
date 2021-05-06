package main
import "regexp"

func main() {
   for _, s := range []string{"o..", "p.."} {
      b := []byte("south north")
      r := regexp.MustCompile(s)
      if r.Match(b) {
         b = r.Find(b)
         println(string(b))
      }
   }
}
