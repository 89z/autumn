package main
import (
   "fmt"
   "github.com/wzshiming/ctc"
)
func main() {
   s1 := "Sunday"
   // example 1
   fmt.Print(ctc.ForegroundGreen)
   fmt.Print(s1)
   fmt.Print(ctc.ForegroundRed)
   fmt.Print(s1)
   fmt.Print(ctc.Reset)
   fmt.Println(s1)
   // example 2
   s2 := ctc.ForegroundGreen.String() + s1
   s3 := ctc.ForegroundRed.String() + s1
   s4 := ctc.Reset.String() + s1
   fmt.Println(s2 + s3 + s4)
}
