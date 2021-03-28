package main
import "golang.org/x/text/cases"

func main() {
   s := "NORTH"
   t := cases.Fold().String(s)
   println(t == "north")
}
