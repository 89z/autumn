package main

import (
   "io"
   "io/ioutil"
   "log"
   "os"
)

type Progress struct {
   Parent io.Reader
   Total int
}

func (o *Progress) Read(y []byte) (int, error) {
   n, e := o.Parent.Read(y)
   o.Total += n
   if n > 0 {
      print("Received ", o.Total / 1000, " K\r")
   } else {
      print("\n")
   }
   return n, e
}

func main() {
   open_o, e := os.Open("10Mio.dat")
   if e != nil {
      log.Fatal(e)
   }
   prog_o := Progress{open_o, 0}
   io.Copy(ioutil.Discard, &prog_o)
}
