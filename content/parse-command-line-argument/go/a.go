package main

import (
   "flag"
   "os"
   "strings"
)

func main() {
   s_con := flag.String("c", "", "contains")
   s_pre := flag.String("p", "", "has prefix")
   flag.Parse()

   if flag.NArg() != 1 {
      println("test [flags] <string>")
      flag.PrintDefaults()
      os.Exit(1)
   }

   s := flag.Arg(0)
   b := strings.HasPrefix(s, *s_pre) && strings.Contains(s, *s_con)
   println(b)
}
