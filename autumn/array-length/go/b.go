package main

func main() {
   { // example 1
      s := []string{"west", "east"}
      n := len(s)
      println(n == 2)
   }
   { // example 2
      s := make(chan string, 9)
      s <- "west"
      n := len(s)
      println(n == 1)
   }
}
