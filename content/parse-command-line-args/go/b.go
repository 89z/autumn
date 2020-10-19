package main

import (
   "flag"
   "os"
)

func main() {
   var n_start, n_len int
   flag.IntVar(&n_start, "a", 0, "start")
   flag.IntVar(&n_len, "b", 1, "length")
   flag.Parse()

   if flag.NArg() != 1 {
      println("slice [flags] <input>")
      flag.PrintDefaults()
      os.Exit(1)
   }

   s_in := flag.Arg(0)
   s_out := s_in[n_start : n_start + n_len]
   println(s_out)
}
