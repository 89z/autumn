package main
import "path/filepath"

func main() {
   path_s := `C:\go\bin\go.exe`
   dir_s := filepath.Dir(path_s)
   println(dir_s == `C:\go\bin`)
}
