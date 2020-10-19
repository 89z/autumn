package main

import (
   "flag"
   "os"
)

func main() {
   n_start := flag.Int("a", 0, "start")
   n_len := flag.Int("b", 1, "length")
   flag.Parse()

   if flag.NArg() != 1 {
      println("slice [flags] <input>")
      flag.PrintDefaults()
      os.Exit(1)
   }

   s_in := flag.Arg(0)
   s_out := s_in[*n_start : *n_start + *n_len]
   println(s_out)
}
