package main
import "github.com/daviddengcn/go-colortext"
func main() {
   b_bold := true
   ct.Foreground(ct.Red, b_bold)
   println("Sunday")
   ct.ResetColor()
   println("Sunday")
}
