package main
import "strconv"

func main() {
   b := strconv.AppendInt(nil, 100, 10)
   println(string(b) == "100")
}
