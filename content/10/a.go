package main
func main() {
   n, s := 1, ""
   switch n {
   case 3:
      s = "all"
   case 2, 1:
      s = "some"
   default:
      s = "none"
   }
   println(s == "some")
}
