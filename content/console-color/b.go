package main
import (
   "github.com/mattn/go-colorable"
   "fmt"
)
func main() {
   s_red := "\x1b[31m"
   s_gre := "\x1b[32m"
   s_end := "\x1b[m"
   s_in := s_red + "Sunday" + s_gre + "Monday" + s_end + "Tuesday"
   o_col := colorable.NewColorableStdout()
   fmt.Fprintln(o_col, s_in)
}
