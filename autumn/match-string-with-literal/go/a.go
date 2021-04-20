package main
import "bytes"

func main() {
   {
      b := bytes.Contains(nil, nil)
      println(b) // true
   }
   {
      b := bytes.Contains([]byte("north"), nil)
      println(b) // true
   }
   {
      b := bytes.Contains([]byte("north"), []byte("or"))
      println(b) // true
   }
}
