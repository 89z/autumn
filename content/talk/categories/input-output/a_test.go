package main

import (
   "bufio"
   "io/ioutil"
   "log"
   "os"
   "regexp"
   "testing"
)

// packages.msys2.org/package/mingw-w64-x86_64-openssl
const s_pac = "mingw-w64-x86_64-openssl"
const s_re = `"(https://repo.msys2.org/[^"]*)"`

// BenchmarkBufio-12          27220             44091 ns/op
func BenchmarkBufio(b *testing.B) {
   o_re := regexp.MustCompile(s_re)
   for n := 0; n < b.N; n++ {
      o_file, e := os.Open(s_pac)
      if e != nil {
         log.Fatal(e)
      }
      o_scan := bufio.NewScanner(o_file)
      for o_scan.Scan() {
         y := o_scan.Bytes()
         a := o_re.FindSubmatch(y)
         if len(a) > 0 {
            break
         }
      }
   }
}

// BenchmarkIoutil-12         11788            101952 ns/op
func BenchmarkIoutil(b *testing.B) {
   o_re := regexp.MustCompile(s_re)
   for n := 0; n < b.N; n++ {
      y_file, e := ioutil.ReadFile(s_pac)
      if e != nil {
         log.Fatal(e)
      }
      o_re.FindSubmatch(y_file)
   }
}
