package main
import "path/filepath"

func main() {
   file := `C:\go\bin\go.exe`
   dir := filepath.Dir(file)
   println(dir == `C:\go\bin`)
}
