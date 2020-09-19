package main
import "path/filepath"

func main() {
   s := `C:\go\bin\go.exe`
   s1 := filepath.Dir(s)
   println(s1 == `C:\go\bin`)
}
