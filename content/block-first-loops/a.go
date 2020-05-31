package main
func main() {
   var n int
   // example 1
   n = 10
   for true {
      println(n)
      if n == 19 {
         break
      }
      n++
   }
   // example 2
   n = 10
   for {
      println(n)
      if n == 19 {
         break
      }
      n++
   }
}
