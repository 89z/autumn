package main

import (
   "encoding/csv"
   "strconv"
   "strings"
)

const s =
"8\t12\tcmd/git-board/{board.go => git-board.go}\n" +
"38\t6\tcmd/youtube-views/{views.go => youtube-views.go}\n"

func main() {
   r := csv.NewReader(strings.NewReader(s))
   r.Comma = '\t'
   for {
      a, e := r.Read()
      if e != nil { break }
      add, e := strconv.Atoi(a[0])
      if e != nil {
         panic(e)
      }
      del, e := strconv.Atoi(a[1])
      if e != nil {
         panic(e)
      }
      println(add, del)
   }
}
