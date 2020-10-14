package main

import (
   "flag"
   "os"
)

func main() {
   n_start := flag.Int("start", 0, "starting position")
   n_len := flag.Int("len", 1, "substring length")
   flag.Parse()

   if flag.NArg() != 1 {
      println("slice [flags] <string>")
      flag.PrintDefaults()
      os.Exit(1)
   }

   s_in := flag.Arg(0)
   s_out := s_in[*n_start : *n_start + *n_len]
   println(s_out)
}
