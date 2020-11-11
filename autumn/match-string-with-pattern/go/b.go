package main
import "regexp"

func Match(sub_s, pat_s string) bool {
   find_a := regexp.MustCompile(pat_s).FindAllString(sub_s, -1)
   return len(find_a) > 0
}

func main() {
   b := Match("January", "a.")
   println(b)
}
