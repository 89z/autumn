package main

func main() {
   println("example 1")
   n := 10
   for true {
      println(n)
      if n == 19 {
         break
      }
      n++
   }
   println("example 2")
   n2 := 10
   for {
      println(n2)
      if n2 == 19 {
         break
      }
      n2++
   }
}
