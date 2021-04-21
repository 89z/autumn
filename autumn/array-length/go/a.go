package main

func main() {
   { // example 1
      a := []string{"west", "east"}
      n := cap(a)
      println(n == 2)
   }
   { // example 1
      a := make(chan string, 2)
      n := cap(a)
      println(n == 2)
   }
}
