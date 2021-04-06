package main
import "os"

func main() {
   e := os.Setenv("MSYSTEM", "MINGW64")
   if e != nil {
      panic(e)
   }
}
