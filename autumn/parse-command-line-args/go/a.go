package main

import (
   "flag"
   "os"
)

func main() {
   s_pre := flag.String("p", "", "prefix")
   s_suf := flag.String("s", "", "suffix")
   flag.Parse()

   if flag.NArg() != 1 {
      println("add [flags] <stem>")
      flag.PrintDefaults()
      os.Exit(1)
   }

   s_stem := flag.Arg(0)
   println(*s_pre + s_stem + *s_suf)
}
