package main
import "os"

func main() {
   e := os.MkdirAll("west/east", os.ModeDir)
   if e != nil {
      panic(e)
   }
}
