package main

import (
   "flag"
   "os"
)

func main() {
   n_len := flag.Int("len", 1, "length")
   n_sta := flag.Int("start", 0, "index")
   flag.Parse()

   if flag.NArg() != 1 {
      println("slice [flags] <string>")
      flag.PrintDefaults()
      os.Exit(1)
   }

   s_in := flag.Arg(0)
   s_out := s_in[*n_sta : *n_sta + *n_len]
   println(s_out)
}
