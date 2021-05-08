package main

import (
   "bufio"
   "fmt"
   "strings"
)

const s =
"8\t12\tcmd/git-board/{board.go => git-board.go}\n" +
"38\t6\tcmd/youtube-views/{views.go => youtube-views.go}\n"

func main() {
   b := bufio.NewScanner(strings.NewReader(s))
   for b.Scan() {
      var add, del int
      fmt.Sscan(b.Text(), &add, &del)
      println(add, del)
   }
}
