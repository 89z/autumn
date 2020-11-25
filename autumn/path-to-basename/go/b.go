package main
import "path/filepath"

func main() {
   in_s := `C:\Windows\notepad.exe`
   out_s := filepath.Base(in_s)
   println(out_s == "notepad.exe")
}
