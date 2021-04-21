package main
import "path/filepath"

func main() {
   s := filepath.Dir(`C:\go\README.md`)
   println(s == `C:\go`)
}
