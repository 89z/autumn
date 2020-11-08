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
   o_in, e := http.Get("http://speedtest.lax.hivelocity.net/10Mio.dat")
   if e != nil {
      log.Fatal(e)
   }
   o_out, e := os.Create(os.DevNull)
   if e != nil {
      log.Fatal(e)
   }
   o_tee := io.TeeReader(o_in.Body, &Progress{})
   io.Copy(o_out, o_tee)
}
