package main

import (
   "flag"
   "os"
)

func main() {
   pre := flag.String("p", "", "prefix")
   suf := flag.String("s", "", "suffix")
   flag.Parse()

   if flag.NArg() != 1 {
      println("add [flags] <stem>")
      flag.PrintDefaults()
      os.Exit(1)
   }

   stem := flag.Arg(0)
   println(*pre + stem + *suf)
}
