package main
import "fmt"

func popString(a *[]string) string {
   n := len(*a)
   v := (*a)[n - 1]
   *a = (*a)[:n - 1]
   return v
}

func main() {
   a := []string{"May", "June", "July"}
   s := popString(&a)
   fmt.Println(a, s)
}
