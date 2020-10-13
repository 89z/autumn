package main
import "flag"

func main() {
   var n int
   flag.IntVar(&n, "year", 0, "print year")
   flag.Parse()
   println(n == 2019)
}
