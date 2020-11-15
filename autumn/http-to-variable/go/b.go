package main

import (
   "io/ioutil"
   "log"
   "net/http"
   "os"
   "path"
)

func IsFile(s string) bool {
   o, e := os.Stat(s)
   return e == nil && o.Mode().IsRegular()
}

func Copy(url_s string) (int64, error) {
   base_s := path.Base(url_s)
   if IsFile(base_s) {
      return 0, nil
   }
   get_o, e := http.Get(url_s)
   if e != nil {
      return 0, e
   }
   create_o, e := os.Create(base_s)
   if e != nil {
      return 0, e
   }
   return create_o.ReadFrom(get_o.Body)
}

func GetContents(url_s string) ([]byte, error) {
   base_s := path.Base(url_s)
   if ! IsFile(base_s) {
      Copy(url_s)
   }
   return ioutil.ReadFile(base_s)
}

func main() {
   y, e := GetContents("http://speedtest.lax.hivelocity.net/index.html")
   if e != nil {
      log.Fatal(e)
   }
   log.Printf("%s", y)
}
