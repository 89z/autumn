package main

func main() {
   n, s := 10, ""
   switch n {
   case 12:
      s = "all"
   case 11, 10:
      s = "some"
   default:
      s = "none"
   }
   println(s == "some")
}
