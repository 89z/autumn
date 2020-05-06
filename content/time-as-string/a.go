package main
import "time"
func main() {
   o1 := time.Now()
   // example 1
   s1 := o1.String()
   // example 2
   s2 := o1.Format(time.ANSIC)
   // example 3
   s3 := o1.Format("Mon Jan 2 15:04:05 2006")
   // print
   print(s1, "\n", s2, "\n", s3, "\n")
}
