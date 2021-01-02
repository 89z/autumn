package main

func main() {
   m := map[int]struct{Month, Day int}{
      2020: {12, 31},
      2019: {11, 30},
   }
   // example 1
   o, b := m[2020]
   if b {
      println(o.Day)
   }
   // example 2
   n := m[2020].Day
   if n != 0 {
      println(n)
   }
}
