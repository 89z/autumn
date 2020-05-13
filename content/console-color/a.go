package main
import "github.com/daviddengcn/go-colortext"
func main() {
   b_bold := false
   ct.Foreground(ct.Red, b_bold)
   print("Sunday")
   ct.Foreground(ct.Green, b_bold)
   print("Monday")
   ct.ResetColor()
   println("Tuesday")
}
