package main

// example 1
func f(d int, e int) int {
   return d + e
}

// example 2
func g(d, e int) int {
   return d + e
}

// print
func main() {
   println(f(4, 5) == 9)
   println(g(4, 5) == 9)
}
