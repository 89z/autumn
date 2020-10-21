package main

import (
   "flag"
   "os"
)

func main() {
   var s_start, s_end string
   flag.StringVar(&s_start, "s", "", "start")
   flag.StringVar(&s_end, "e", "", "end")
   flag.Parse()

   if flag.NArg() != 1 {
      println("cat [flags] <input>")
      flag.PrintDefaults()
      os.Exit(1)
   }

   s_in := flag.Arg(0)
   println(s_start + s_in + s_end)
}
