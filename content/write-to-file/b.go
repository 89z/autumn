package main
import (
   "bytes"
   "os"
)
func main() {
   var f_in bytes.Buffer
   f_in.WriteString("Sunday\n")
   var f_out, _ = os.Create("a.txt")
   f_in.WriteTo(f_out)
}
