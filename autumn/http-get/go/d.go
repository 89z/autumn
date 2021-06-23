package main

import (
   "fmt"
   "io"
   "net/http"
)

type body struct {
   io.Reader
   read int
}

func (b *body) Read(p []byte) (int, error) {
   n, err := b.Reader.Read(p)
   if err != nil {
      fmt.Println()
   } else {
      b.read += n
      fmt.Printf("\rRead %9v", b.read)
   }
   return n, err
}

func main() {
   r, err := http.Get("http://speedtest.lax.hivelocity.net/10Mio.dat")
   if err != nil {
      panic(err)
   }
   defer r.Body.Close()
   io.ReadAll(&body{Reader: r.Body})
}
