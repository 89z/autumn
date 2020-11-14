package main

import (
   "bytes"
   "io"
   "os"
   "strings"
   "testing"
)

func Bytes(path_s string) (string, error) {
   open_o, e := os.Open(path_s)
   if e != nil {
      return "", e
   }
   defer open_o.Close()
   buf_o := bytes.Buffer{}
   buf_o.ReadFrom(open_o)
   return buf_o.String(), nil
}

func Strings(path_s string) (string, error) {
   open_o, e := os.Open(path_s)
   if e != nil {
      return "", e
   }
   defer open_o.Close()
   str_o := strings.Builder{}
   io.Copy(&str_o, open_o)
   return str_o.String(), nil
}

func BenchmarkBytes(b *testing.B) {
   for n := 0; n < b.N; n++ {
      Bytes("10Mio.dat")
   }
}

func BenchmarkStrings(b *testing.B) {
   for n := 0; n < b.N; n++ {
      Strings("10Mio.dat")
   }
}
