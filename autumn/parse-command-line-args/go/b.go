package main

import (
   "flag"
   "os"
)

func main() {
   var pre_s, suf_s string
   flag.StringVar(&pre_s, "p", "", "prefix")
   flag.StringVar(&suf_s, "s", "", "suffix")
   flag.Parse()

   if flag.NArg() != 1 {
      println("add [flags] <stem>")
      flag.PrintDefaults()
      os.Exit(1)
   }

   stem_s := flag.Arg(0)
   println(pre_s + stem_s + suf_s)
}
