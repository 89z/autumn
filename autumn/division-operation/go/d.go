package main

func main() {
   { // natural int
      n := 7 / 2
      println(n == 3)
   }
   { // natural float
      n := 7.0 / 2
      println(n == 3.5)
   }
   { // force int
      n := int(7.0) / 2
      println(n == 3)
   }
   { // force float
      n := float64(7) / 2
      println(n == 3.5)
   }
}
