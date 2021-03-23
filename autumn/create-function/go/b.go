package main

func main() {
   { // example 1
      f := func (d int, e int) int { return d + e }
      println(f(4, 5) == 9)
   }
   { // example 2
      f := func (d, e int) int { return d + e }
      println(f(4, 5) == 9)
   }
}
