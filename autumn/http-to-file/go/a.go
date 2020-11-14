package main

import (
   "io"
   "log"
   "net/http"
   "os"
)

type Progress struct {
   Received int
}

func (o *Progress) Write(y []byte) (int, error) {
   n := len(y)
   o.Received += n
   println(o.Received)
   return n, nil
}

func main() {
   get_o, e := http.Get("http://speedtest.lax.hivelocity.net/10Mio.dat")
   if e != nil {
      log.Fatal(e)
   }
   create_o, e := os.Create(os.DevNull)
   if e != nil {
      log.Fatal(e)
   }
   tee_o := io.TeeReader(get_o.Body, &Progress{})
   io.Copy(create_o, tee_o)
}
