package main

type date struct { month, day int }

func newDate() date {
   return date{1, 1}
}

func (d *date) add() { d.day++ }

func main() {
   d := newDate()
   d.add()
   println(d.day == 2)
}
