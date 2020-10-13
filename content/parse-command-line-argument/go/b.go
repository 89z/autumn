package main
import "flag"

func main() {
   var b bool
   flag.BoolVar(&b, "year", false, "print year")
   flag.Parse()
   println(b)
}
