package main
import "path/filepath"

func main() {
   s := `C:\go\bin\go.exe`
   s2 := filepath.Base(s)
   println(s2 == "go.exe")
}
