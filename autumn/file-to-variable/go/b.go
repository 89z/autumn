package main
import "os"

func main() {
   b, e := os.ReadFile("a.go")
   if e != nil {
      panic(e)
   }
   os.Stdout.Write(b)
}
