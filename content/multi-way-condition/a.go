package main
func main() {
   var s = "Sun"
   var n int
   switch s {
   case "Fri":
      n = 1
   case "Sat", "Sun":
      n = 2
   default:
      n = 0
   }
   println(n)
}
