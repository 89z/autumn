package main
import (
   "fmt"
   "os"
   "time"
)
func main() {
   o1 := time.Now()
   os.Chtimes("index.md", o1, o1)
   o2, e := os.Stat("index.md")
   if e != nil {
      os.Exit(1)
   }
   n1 := o1.Unix()
   n2 := o2.ModTime().Unix()
   fmt.Println(n1, n2 == n1)
}
