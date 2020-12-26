package main

func main() {
   m := map[int]struct{Month, Day int}{
      2019: {12, 31},
      2018: {11, 31},
   }
   // example 1
   o, b := m[2019]
   if b {
      println(o.Day)
   }
   // example 2
   n := m[2019].Day
   if n != 0 {
      println(n)
   }
}
