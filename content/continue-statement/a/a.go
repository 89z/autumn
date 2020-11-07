package main

func main() {
   for n := 0; n < 10; n++ {
      if n % 3 == 0 {
         continue
      }
      println(n)
   }
}
