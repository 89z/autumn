package main

import (
   "fmt"
   "io"
   "io/ioutil"
   "log"
   "math"
   "net/http"
)

func NumberFormat(n float64) string {
   n2 := int(math.Log10(n)) / 3
   n3 := n / math.Pow10(n2 * 3)
   return fmt.Sprintf("%.3f", n3) + []string{"", " k", " M", " G"}[n2]
}

type Progress struct {
   Parent io.Reader
   Total float64
}

func (o *Progress) Read(y []byte) (int, error) {
   n, e := o.Parent.Read(y)
   if e != nil {
      fmt.Println()
   } else {
      o.Total += float64(n)
      fmt.Printf("READ %9s\r", NumberFormat(o.Total))
   }
   return n, e
}

func main() {
   get_o, e := http.Get("http://speedtest.lax.hivelocity.net/10Mio.dat")
   if e != nil {
      log.Fatal(e)
   }
   prog_o := Progress{get_o.Body, 0}
   io.Copy(ioutil.Discard, &prog_o)
}
