package main
import "path/filepath"

func main() {
   s := `C:\go\bin\go.exe`
   s1 := filepath.Base(s)
   println(s1 == "go.exe")
}
