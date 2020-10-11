package main

import (
   "bufio"
   "fmt"
   "log"
   "net/http"
   "regexp"
)

func main() {
   s := "https://packages.msys2.org/package/mingw-w64-x86_64-openssl"
   o_get, e := http.Get(s)
   if e != nil {
      log.Fatal(e)
   }
   o_scan := bufio.NewScanner(o_get.Body)
   o_re := regexp.MustCompile(`"(https://repo.msys2.org/[^"]*)"`)
   for o_scan.Scan() {
      y := o_scan.Bytes()
      a := o_re.FindSubmatch(y)
      if len(a) > 0 {
         y := a[1]
         fmt.Printf("%s\n", y)
         break
      }
   }
}
