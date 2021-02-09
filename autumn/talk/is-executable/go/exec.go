//go:generate mkwinsyscall -output zexec.go exec.go
//sys GetBinaryType(name string, binaryType *int) (err error) = kernel32.GetBinaryTypeW
package main
import "log"

func main() {
   var n int
   e := GetBinaryType("exec.exe", &n)
   if e != nil {
      log.Fatal(e)
   }
   println(n)
}
