package main
import "bytes"
func main() {
   a1 := []byte{'S', 'u', 'n'}
   s1 := bytes.NewBuffer(a1).String()
   println(s1)
}
