package main

import (
   "bufio"
   "log"
   "os"
)

func File(path_s string) ([]string, error) {
   a := []string{}
   open_o, e := os.Open(path_s)
   if e != nil {
      return a, e
   }
   scan_o := bufio.NewScanner(open_o)
   for scan_o.Scan() {
      text_s := scan_o.Text()
      a = append(a, text_s)
   }
   return a, nil
}

func main() {
   a, e := File("a.go")
   if e != nil {
      log.Fatal(e)
   }
   log.Printf("%q", a)
}
