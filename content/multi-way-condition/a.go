package main
func main() {
   n := 0
   s := "Sunday"
   switch s {
   case "Friday":
      n = 2
   case "Saturday", "Sunday":
      n = 3
   default:
      n = 1
   }
   println(n)
}
