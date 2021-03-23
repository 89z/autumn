package main

import (
   "flag"
   "os"
)

func main() {
   var pre, suf string
   flag.StringVar(&pre, "p", "", "prefix")
   flag.StringVar(&suf, "s", "", "suffix")
   flag.Parse()

   if flag.NArg() != 1 {
      println("add [flags] <stem>")
      flag.PrintDefaults()
      os.Exit(1)
   }

   stem := flag.Arg(0)
   println(pre + stem + suf)
}
