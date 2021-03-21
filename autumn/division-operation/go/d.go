package main

func main() {
   f, i := 7.5, 7
   { // example 1
      n := f / 2
      println(n == 3.75)
   }
   { // example 2
      n := int(f) / 2
      println(n == 3)
   }
   { // example 3
      n := i / 2
      println(n == 3)
   }
   { // example 4
      n := float64(i) / 2
      println(n == 3.5)
   }
}
