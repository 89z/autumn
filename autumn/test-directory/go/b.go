package main
import "os"

func main() {
   o, e := os.Stat(`C:\Users`)
   if e != nil || ! o.IsDir() {
      os.Exit(1)
   }
}
