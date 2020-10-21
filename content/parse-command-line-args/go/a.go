package main

import (
   "flag"
   "os"
)

func main() {
   s_start := flag.String("s", "", "start")
   s_end := flag.String("e", "", "end")
   flag.Parse()

   if flag.NArg() != 1 {
      println("cat [flags] <input>")
      flag.PrintDefaults()
      os.Exit(1)
   }

   s_in := flag.Arg(0)
   println(*s_start + s_in + *s_end)
}
