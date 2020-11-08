package main

func main() {
   var n = 1 + 2
   var b bool

   switch n {
   case 5:
      b = false
   case 4, 3:
      b = true
   default:
      b = n < 3
   }

   println(b)
}
