package main
import "path/filepath"
func main() {
   s1 := `C:\go\bin\go.exe`
   s2 := filepath.Base(s1)
   println(s2 == "go.exe")
}
