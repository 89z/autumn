package main
import "time"
func main() {
   o1 := time.Now()
   // example 1
   s1 := o1.String()
   // example 2
   s2 := o1.Format(time.ANSIC)
   // example 3
   s3 := o1.Format("Mon Jan 2 2006 15:04:05")
   // example 4
   a1, _ := o1.MarshalText()
   s4 := string(a1)
   // print
   print(s1, "\n", s2, "\n", s3, "\n", s4, "\n")
}
