package main
import "path/filepath"
func main() {
   s1 := `C:\go\bin\go.exe`
   s2 := filepath.Dir(s1)
   println(s2 == `C:\go\bin`)
}
