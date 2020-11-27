package main
type Num int

func (n Num) add(n2 Num) Num {
   return n + n2
}

func main() {
   n1 := Num(1).add(2).add(3)
   n2 := Num.add(1, 2).add(3)
   println(n1 == 6, n2 == 6)
}
