package main
import "bytes"

func main() {
   {
      b := bytes.Contains([]byte{}, []byte{})
      println(b) // true
   }
   {
      b := bytes.Contains([]byte("north"), []byte{})
      println(b) // true
   }
   {
      b := bytes.Contains([]byte("north"), []byte("or"))
      println(b) // true
   }
}
