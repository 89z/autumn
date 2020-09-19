package main

func main() {
   s := "📗"
   // example 1
   n1 := len(s)
   // example 2
   a := []rune(s)
   n2 := len(a)
   // print
   println(n1 == 4, n2 == 1)
}
