package main
import "os"

func main() {
   s := "March\n"
   o, e := os.Create("a.txt")
   if e != nil {
      os.Exit(1)
   }
   o.WriteString(s)
}
