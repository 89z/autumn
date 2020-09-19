package main

func main() {
   t := map[string]bool{"May": true, "June": true}
   // example 1
   b1, b1a := t["May"]
   // example 2
   b2 := t["May"]
   // print
   println(b1, b1a, b2)
}
