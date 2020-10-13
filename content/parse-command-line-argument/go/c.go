package main
import "flag"

func main() {
   n := flag.Int("year", 0, "print year")
   flag.Parse()
   println(*n == 2019)
}
