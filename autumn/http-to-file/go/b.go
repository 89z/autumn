package main

import (
   "fmt"
   "io"
   "net/http"
   "os"
)

type progress struct {
   io.Reader
   total int
}

func (p *progress) Read(b []byte) (int, error) {
   n, e := p.Reader.Read(b)
   if e != nil {
      fmt.Println()
   } else {
      p.total += n
      fmt.Printf("\rRead %9v", p.total)
   }
   return n, e
}

func httpCopy(source, dest string) (int64, error) {
   get, e := http.Get(source)
   if e != nil {
      return 0, e
   }
   create, e := os.Create(dest)
   if e != nil {
      return 0, e
   }
   defer create.Close()
   return io.Copy(create, &progress{Reader: get.Body})
}

func main() {
   httpCopy("http://speedtest.lax.hivelocity.net/10Mio.dat", "10Mio.dat")
}
