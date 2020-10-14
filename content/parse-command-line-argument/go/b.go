package main

import (
   "flag"
   "os"
)

func main() {
   var n_len, n_sta int
   flag.IntVar(&n_len, "len", 1, "length")
   flag.IntVar(&n_sta, "start", 0, "index")
   flag.Parse()

   if flag.NArg() != 1 {
      println("slice [flags] <string>")
      flag.PrintDefaults()
      os.Exit(1)
   }

   s_in := flag.Arg(0)
   s_out := s_in[n_sta : n_sta + n_len]
   println(s_out)
}
