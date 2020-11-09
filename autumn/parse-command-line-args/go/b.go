package main

import (
   "flag"
   "os"
)

func main() {
   var s_pre, s_suf string
   flag.StringVar(&s_pre, "p", "", "prefix")
   flag.StringVar(&s_suf, "s", "", "suffix")
   flag.Parse()

   if flag.NArg() != 1 {
      println("add [flags] <stem>")
      flag.PrintDefaults()
      os.Exit(1)
   }

   s_stem := flag.Arg(0)
   println(s_pre + s_stem + s_suf)
}
