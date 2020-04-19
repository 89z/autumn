package main
import (
   "compress/gzip"
   "io"
   "os"
)
func main() {
   // head
   o1, _ := os.Open("a.tar.gz")
   o2, _ := gzip.NewReader(o1)
   o3, _ := os.Create("a.tar")
   // body
   io.Copy(o3, o2)
}
