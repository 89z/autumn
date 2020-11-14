package main

import (
   "io"
   "io/ioutil"
   "net/http"
   "os"
   "path"
   "testing"
)

func BenchmarkReadFrom(b *testing.B) {
   for n := 0; n < b.N; n++ {
      ReadFrom("https://packages.msys2.org/package/mingw-w64-x86_64-gcc")
   }
}

func ReadFrom(url_s string) ([]byte, error) {
   get_o, e := http.Get(url_s)
   if e != nil {
      return []byte{}, e
   }
   base_s := path.Base(url_s)
   create_o, e := os.Create(base_s)
   if e != nil {
      return []byte{}, e
   }
   create_o.ReadFrom(get_o.Body)
   return ioutil.ReadFile(base_s)
}

func BenchmarkTeeReader(b *testing.B) {
   for n := 0; n < b.N; n++ {
      TeeReader("https://packages.msys2.org/package/mingw-w64-x86_64-gcc")
   }
}

func TeeReader(url_s string) ([]byte, error) {
   get_o, e := http.Get(url_s)
   if e != nil {
      return []byte{}, e
   }
   base_s := path.Base(url_s)
   create_o, e := os.Create(base_s)
   if e != nil {
      return []byte{}, e
   }
   tee_o := io.TeeReader(get_o.Body, create_o)
   return ioutil.ReadAll(tee_o)
}
