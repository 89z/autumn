package main

type date struct { month, day int }

func newDate() date {
   return date{12, 31}
}

func (d *date) year() {
   d.month = 1
   d.day = 1
}

func main() {
   d := newDate()
   d.year()
   println(d.day == 1)
}
