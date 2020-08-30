package main
import "path/filepath"
func main() {
   s1 := "a.go"
   s2 := filepath.Ext(s1)
   println(s2 == ".go")
}
