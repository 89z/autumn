package main

func main() {
   var s = "minute"
   var n int

   switch s {
   case "hour":
      n = 23
   case "minute", "second":
      n = 59
   default:
      n = -1
   }

   println(n == 59)
}
