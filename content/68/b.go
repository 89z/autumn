package main
import "path/filepath"

func main() {
   s := `C:\go\bin\go.exe`
   s2 := filepath.Dir(s)
   println(s2 == `C:\go\bin`)
}
