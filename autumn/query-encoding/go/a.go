package main

import (
   "fmt"
   "net/url"
)

func main() {
   { // example 1
      q, e := url.ParseQuery("west=left&east=right")
      if e != nil {
         panic(e)
      }
      fmt.Println(q) // map[east:[right] west:[left]]
   }
   { // example 2
      q, e := url.ParseQuery("http://youtube.com/watch?v=GvvRUKKXilg")
      if e != nil {
         panic(e)
      }
      fmt.Println(q) // map[http://youtube.com/watch?v:[GvvRUKKXilg]]
   }
}
