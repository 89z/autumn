package main

func main() {
   { // example 1
      a := []string{"west", "east"}
      n := len(a)
      println(n == 2)
   }
   { // example 2
      a := make(chan string, 9)
      a <- "west"
      n := len(a)
      println(n == 1)
   }
}
