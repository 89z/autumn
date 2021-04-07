package main
import "os"

func main() {
   e := os.MkdirAll("north", os.ModeDir)
   if e != nil {
      panic(e)
   }
}
