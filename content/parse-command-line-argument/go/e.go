package main
import "flag"

func main() {
   s := flag.String("month", "", "print month")
   flag.Parse()
   println(*s == "May")
}
