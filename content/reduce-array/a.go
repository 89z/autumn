package main
import "fmt"

func main() {
   a := []string{"May", "June"}
   m := map[string]int{}
   for n, s := range a {
      m[s] = n
   }
   fmt.Println(m)
}
