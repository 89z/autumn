package main

func main() {
   s := "📗"
   // example 1
   n := len(s)
   // example 2
   a := []rune(s)
   n2 := len(a)
   // print
   println(n == 4, n2 == 1)
}
