package main

import (
   "bytes"
   "io"
   "net/http"
   "os"
   "path"
)

func Get(url_s string, ret_b bool) (string, error) {
   base_s := path.Base(url_s)
   get_o, e := http.Get(url_s)
   if e != nil {
      return "", e
   }
   create_o, e := os.Create(base_s)
   if e != nil {
      return "", e
   }
   if ! ret_b {
      io.Copy(create_o, get_o.Body)
      return "", nil
   }
   buf_o := bytes.Buffer{}
   tee_o := io.TeeReader(get_o.Body, &buf_o)
   io.Copy(create_o, tee_o)
   return buf_o.String(), nil
}

func main() {
   Get("https://speedtest.lax.hivelocity.net/index.html", false)
}
