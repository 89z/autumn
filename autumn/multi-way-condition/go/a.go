package main

func main() {
   var (
      n = 0x63 - 0x20
      r rune
   )

   switch n {
   case 0x41:
      r = 'A'
   case 0x42, 0x62:
      r = 'B'
   default:
      r = rune(n)
   }

   println(r == 'C')
}
