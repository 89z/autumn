package main
import "regexp"

func Match(SubS, PatS string) bool {
   FindA := regexp.MustCompile(PatS).FindAllString(SubS, -1)
   return len(FindA) > 0
}

func main() {
   b := Match("January", "a.")
   println(b)
}
