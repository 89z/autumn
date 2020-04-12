package main
func main() {
   // example 1
   if true {
      const n1 = 10
   }
   println(n1) // undefined: n1
   // example 2
   switch true {
   case true:
      const n2 = 10
   }
   println(n2) // undefined: n2
   // example 3
   // map[bool]int literal[true] is not a constant
   const n3 = map[bool]int{true: 10}[true]
   // example 4
   m1 := map[bool]int{true: 10}
   const n4 = m1[true] // m1[true] is not a constant
   // example 5
   const m2 = map[bool]int{} // map[bool]int literal is not a constant
   // example 6
   // make(map[bool]int) is not a constant
   const m3 = make(map[bool]int)
   // example 7
   // (func literal)() is not a constant
   const n5 = func() int { return 10 }()
   // example 8
   u1 := func() int { return 10 }
   const n6 = u1() // u1() is not a constant
   // example 9
   // func literal is not a constant
   const u2 = func() int { return 10 }
}
