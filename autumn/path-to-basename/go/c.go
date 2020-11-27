package main
import "strings"

func BaseName(s string, y byte) string {
   n := strings.LastIndexByte(s, y)
   if n == -1 {
      return s
   }
   return s[:n]
}

func main() {
   s := BaseName(`C:\Windows\notepad.exe`, '.')
   println(s == `C:\Windows\notepad`)
}
