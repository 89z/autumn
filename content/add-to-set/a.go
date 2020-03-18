package main
import "fmt"
func main() {
   c1 := map[string]bool{"Sun": true}
   c1["Mon"] = true
   fmt.Println(c1)
}
