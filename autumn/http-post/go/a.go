package main

import (
   "bytes"
   "io"
   "mime/multipart"
   "net/http"
   "os"
   "strings"
)

func createForm(form map[string]string) (string, io.Reader, error) {
   body := new(bytes.Buffer)
   mp := multipart.NewWriter(body)
   defer mp.Close()
   for key, val := range form {
      if strings.HasPrefix(val, "@") {
         val = val[1:]
         file, err := os.Open(val)
         if err != nil { return "", nil, err }
         defer file.Close()
         part, err := mp.CreateFormFile(key, val)
         if err != nil { return "", nil, err }
         io.Copy(part, file)
      } else {
         mp.WriteField(key, val)
      }
   }
   return mp.FormDataContentType(), body, nil
}

func main() {
   form := map[string]string{"file": "@a.go", "account": "5555555555"}
   ct, body, err := createForm(form)
   if err != nil {
      panic(err)
   }
   http.Post("https://github.com", ct, body)
}
