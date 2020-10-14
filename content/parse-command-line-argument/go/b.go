package main

import (
   "flag"
   "os"
   "strings"
)

func main() {
   var s_pre, s_con string
   flag.StringVar(&s_pre, "p", "", "has prefix")
   flag.StringVar(&s_con, "c", "", "contains")
   flag.Parse()

   if flag.NArg() != 1 {
      println("test [flags] <string>")
      flag.PrintDefaults()
      os.Exit(1)
   }

   s := flag.Arg(0)
   b := strings.HasPrefix(s, s_pre) && strings.Contains(s, s_con)
   println(b)
}
