package main
import "path/filepath"

func main() {
   s, t := filepath.Split(`C:\go\README.md`)
   println(s == `C:\go\`, t == "README.md")
}
