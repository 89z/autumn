package main

func main() {
   println("example 1")
   n1 := 10
   for n1 < 20 {
      println(n1)
      n1++
   }

   println("example 2")
   n2 := 10
   for {
      if n2 > 19 {
         break
      }
      println(n2)
      n2++
   }
}
