package main
import "os"

func main() {
   s, e := os.UserCacheDir()
   if e != nil {
      panic(e)
   }
   println(s == `C:\Users\Steven\AppData\Local`)
}
