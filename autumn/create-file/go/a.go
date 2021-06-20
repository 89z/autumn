package main
import "os"

func main() {
   f, e := os.OpenFile("file.txt", os.O_CREATE|os.O_EXCL, os.ModePerm)
   if e != nil {
      panic(e)
   }
   f.Close()
}
