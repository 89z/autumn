package main
import "net/url"

func main() {
   q, e := url.ParseQuery("west=left&east=right")
   if e != nil {
      panic(e)
   }
   s := q.Get("west")
   println(s == "left")
}
