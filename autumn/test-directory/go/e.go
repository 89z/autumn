package main
import "os"

func main() {
   o, e := os.OpenFile("a.txt", os.O_WRONLY|os.O_CREATE|os.O_EXCL, os.ModePerm)
   if e != nil {
      os.Exit(1)
   }
   o.WriteString("March\n")
}
