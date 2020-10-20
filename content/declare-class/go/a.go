package main

type Time struct {
   Hours int
}

func (o Time) Duration(minutes int) int {
   return o.Hours * 60 + minutes
}

func main() {
   o := Time{23}
   n := o.Duration(59)
   println(n == 1439)
}
