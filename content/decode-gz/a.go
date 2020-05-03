package main
import (
   "compress/gzip"
   "io"
   "os"
)
func main() {
   o1, _ := os.Open("a.tar.gz")
   o2, _ := gzip.NewReader(o1)
   o3, _ := os.Create("a.tar")
   // begin
   io.Copy(o3, o2)
}
