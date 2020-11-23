package main
import "path/filepath"

func main() {
   s := `C:\Windows\notepad.exe`
   s1 := filepath.Dir(s)
   println(s1 == `C:\Windows`)
}
