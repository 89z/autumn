package main

import (
   "crypto/aes"
   "crypto/cipher"
   "encoding/base64"
   "log"
   "bytes"
)

func PKCS5Padding(src []byte, blockSize int) []byte {
   padding := blockSize - len(src)%blockSize
   padtext := bytes.Repeat([]byte{byte(padding)}, padding)
   return append(src, padtext...)
}

func main() {
   key, iv := []byte("KDKDKDKDKDKDKDKD"), []byte("IVIVIVIVIVIVIVIV")
   block, e := aes.NewCipher(key)
   if e != nil {
      log.Fatal(e)
   }
   pt := []byte("January February")
   pt = PKCS5Padding(pt, 16)
   cipher.NewCBCEncrypter(block, iv).CryptBlocks(pt, pt)
   ct := base64.StdEncoding.EncodeToString(pt)
   println(ct == "BvfnZp4jmCaveE6kefhumpZ0raWX9GDojfPasgSwLTM=")
}