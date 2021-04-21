package main
import "path/filepath"

func main() {
   s := filepath.Base(`C:\go\README.md`)
   println(s == "README.md")
}
