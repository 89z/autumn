package main
import "flag"

func main() {
   b := flag.Bool("year", false, "print year")
   flag.Parse()
   println(*b)
}
