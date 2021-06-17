package main
import "strings"

func scan(r *strings.Reader, sep byte) ([]byte, error) {
   var b []byte
   for {
      c, e := r.ReadByte()
      if c == sep && b != nil {
         break
      }
      if e != nil && b != nil {
         break
      }
      if e != nil && b == nil {
         return nil, e
      }
      b = append(b, c)
   }
   return b, nil
}

func main() {
   r := strings.NewReader("north,east,south,west")
   for {
      b, e := scan(r, ',')
      if e != nil {
         break
      }
      println(string(b))
   }
}
