package main

import (
   "encoding/json"
   "net/http"
)

func JsonGetReader(r io.Reader) (Map, error) {
   v := Map{}
   return v, json.NewDecoder(r).Decode(&v)
}
