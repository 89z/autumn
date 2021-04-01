package main
import "path/filepath"

func main() {
   s := filepath.Ext("a.go")
   println(s == ".go")
}
