package main
import "os"

func main() {
   s, e := os.UserCacheDir()
   if e != nil {
      os.Exit(1)
   }
   println(s == `C:\Users\Steven\AppData\Local`)
}
