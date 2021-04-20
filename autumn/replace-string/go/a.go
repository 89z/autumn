package main
import "regexp"

func main() {
   s := "west? east.txt"
   s = regexp.MustCompile(`["*/:<>?\|]`).ReplaceAllString(s, "")
   println(s == "west east.txt")
}
