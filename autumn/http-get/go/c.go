package main

import (
   "bufio"
   "net/http"
   "strings"
)

func readRequest(raw, scheme string) (*http.Request, error) {
   r, err := http.ReadRequest(bufio.NewReader(strings.NewReader(raw)))
   if err != nil { return nil, err }
   r.RequestURI, r.URL.Scheme, r.URL.Host = "", scheme, r.Host
   return r, nil
}

func main() {
   raw := `GET /images HTTP/1.1
Host: example.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif
Accept-Encoding: gzip, deflate
Connection: close

`
   req, err := readRequest(raw, "http")
   if err != nil {
      panic(err)
   }
   new(http.Client).Do(req)
}
