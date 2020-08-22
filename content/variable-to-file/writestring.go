package main
import "os"
func main() {
   o, e := os.Create("a.txt")
   if e != nil {
      os.Exit(1)
   }
   o.WriteString("Sunday\n")
}
