package main

import (
   "fmt"
   "io"
   "net/http"
   "os"
)

type Progress struct {
   Parent io.Reader
   Total int
}

func (o *Progress) Read(y []byte) (int, error) {
   n, e := o.Parent.Read(y)
   if e != nil {
      fmt.Println()
   } else {
      o.Total += n
      fmt.Printf("\rRead %9v", o.Total)
   }
   return n, e
}

func HttpCopy(source, dest string) (int64, error) {
   get_o, e := http.Get(source)
   if e != nil {
      return 0, e
   }
   create_o, e := os.Create(dest)
   if e != nil {
      return 0, e
   }
   defer create_o.Close()
   return io.Copy(create_o, &Progress{Parent: get_o.Body})
}

func main() {
   HttpCopy("http://speedtest.lax.hivelocity.net/10Mio.dat", "10Mio.dat")
}
