package main

type Time struct {
   Hours int
}

func NewTime() Time {
   return Time{23}
}

func (o Time) Duration(minutes int) int {
   return o.Hours * 60 + minutes
}

func main() {
   o := NewTime()
   n := o.Duration(59)
   println(n == 1439)
}
