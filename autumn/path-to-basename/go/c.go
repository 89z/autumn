package main
import "strings"

func File(s string) string {
   n := strings.LastIndexByte(s, '.')
   return s[:n]
}

func main() {
   s := File(`C:\Windows\notepad.exe`)
   println(s == `C:\Windows\notepad`)
}
