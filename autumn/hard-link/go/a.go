package main
import "os"

func main() {
   e := os.Link("a.go", "b.go")
   if e != nil {
      panic(e)
   }
}
