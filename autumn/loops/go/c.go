package main

func main() {
   var a [12]string
   // example 1
   for n := range a {
      println(n)
   }
   // example 2
   for range a {
      println("north")
   }
}
