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
   s := BaseName(`C:\go\bin\go.exe`, '.')
   println(s == `C:\go\bin\go`)
}
