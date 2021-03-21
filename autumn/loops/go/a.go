package main

func main() {
   var n int
   println("example 1")
   n = 10
   for n < 20 {
      println(n)
      n++
   }
   println("example 2")
   n = 10
   for {
      if n > 19 {
         break
      }
      println(n)
      n++
   }
}
