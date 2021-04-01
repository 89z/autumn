package main
import "path/filepath"

func main() {
   s := filepath.Ext(`C:\go\README.md`)
   println(s == ".md")
}
