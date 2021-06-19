package main
import "os"

func main() {
   e := os.Remove("file.txt")
   if e != nil {
      panic(e)
   }
}
