package main
func main() {
   // example 1
   n1 := 10
   for true {
      println(n1)
      if n1 == 19 {
         break
      }
      n1++
   }
   // example 2
   n2 := 10
   for {
      println(n2)
      if n2 == 19 {
         break
      }
      n2++
   }
}
