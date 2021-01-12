package main
import "golang.org/x/text/cases"

func main() {
   s := "March"
   t := cases.Fold().String(s)
   println(t == "march")
}
