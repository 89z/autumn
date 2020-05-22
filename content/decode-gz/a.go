package main
import (
   "compress/gzip"
   "io"
   "log"
   "os"
)
func main() {
   o1, e := os.Open("a.tar.gz")
   if e != nil {
      log.Fatal(e)
   }
   o2, e := gzip.NewReader(o1)
   if e != nil {
      log.Fatal(e)
   }
   o3, e := os.Create("a.tar")
   if e != nil {
      log.Fatal(e)
   }
   io.Copy(o3, o2)
}
