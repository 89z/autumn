package main
import "os"

func main() {
   a := os.Environ()
   for _, s := range a {
      println(s)
   }
}
