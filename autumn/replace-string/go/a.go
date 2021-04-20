package main
import "regexp"

func main() {
   s := "north?.txt"
   t := regexp.MustCompile(`["*/:<>?\|]`).ReplaceAllString(s, "+")
   println(t == "north+.txt")
}
