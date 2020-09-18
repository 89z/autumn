package main

func main() {
   t := map[string]bool{"May": true}
   // example A
   bAa, bAb := t["May"]
   // example B
   bB := t["May"]
   // print
   println(bAa, bAb, bB)
}
