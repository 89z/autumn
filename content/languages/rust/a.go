package main
import "fmt"
func main() {
   type Pref struct {
      k string
      n int
      b bool
   }
   o1 := Pref { k: "Sun", n: 10 }
   fmt.Printf("%+v\n", o1)
}
