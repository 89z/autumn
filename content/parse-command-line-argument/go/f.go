package main
import "flag"

func main() {
   var s string
   flag.StringVar(&s, "month", "", "print month")
   flag.Parse()
   println(s == "May")
}
