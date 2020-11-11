package main
import "os"

func main() {
   o, e := os.Stat(`C:\Users`)
   if os.IsNotExist(e) || ! o.IsDir() {
      os.Exit(1)
   }
}
