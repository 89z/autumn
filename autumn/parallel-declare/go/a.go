package main
import "fmt"

func main() {
   { // example 1
      s, n := "year", 2021
      fmt.Println(s, n)
   }
   { // example 2
      s, n := []byte("month"), int64(12)
      fmt.Println(s, n)
   }
}
