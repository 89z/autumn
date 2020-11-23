package main
import "path/filepath"

func main() {
   s := `C:\Windows\notepad.exe`
   s1 := filepath.Base(s)
   println(s1 == "notepad.exe")
}
