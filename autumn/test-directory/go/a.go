package main
import "os"

func main() {
   o, e := os.Stat(`C:\Users`)
   if e != nil || ! o.Mode().IsDir() {
      os.Exit(1)
   }
}
