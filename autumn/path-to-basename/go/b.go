package main
import "path/filepath"

func main() {
   in_s := `C:\go\bin\go.exe`
   out_s := filepath.Base(in_s)
   println(out_s == "go.exe")
}
