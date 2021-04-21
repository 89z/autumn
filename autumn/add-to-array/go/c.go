package main

func main() {
   a := make(chan int, 2)
   a <- 10
   println(<-a)
}
