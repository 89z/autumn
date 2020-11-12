package main

import (
   "flag"
   "os"
)

func main() {
   pre_s := flag.String("p", "", "prefix")
   suf_s := flag.String("s", "", "suffix")
   flag.Parse()

   if flag.NArg() != 1 {
      println("add [flags] <stem>")
      flag.PrintDefaults()
      os.Exit(1)
   }

   stem_s := flag.Arg(0)
   println(*pre_s + stem_s + *suf_s)
}
