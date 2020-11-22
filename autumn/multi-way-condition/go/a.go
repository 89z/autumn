package main

func main() {
   var (
      r = '\x43'
      n int
   )

   switch r {
   case 'A':
      n = 0x41
   case 'B', 'b':
      n = 0x42
   default:
      n = int(r)
   }

   println(n == 0x43)
}
