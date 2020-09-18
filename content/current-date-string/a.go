package main
import "time"

func main() {
   o := time.Now()
   // example 1
   s1 := o.String()
   // example 2
   s2 := o.Format(time.RFC3339)
   // print
   print(s1, "\n", s2, "\n")
}
