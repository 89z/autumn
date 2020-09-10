package main

func main() {
   t := map[string]bool{"May": true}
   // example 1
   b, b2 := t["May"]
   // example 2
   b3 := t["May"]
   // print
   println(b, b2, b3)
}
