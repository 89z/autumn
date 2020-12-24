package main
import "golang.org/x/text/cases"

func main() {
   s := cases.Fold().String("March")
   println(s == "march")
}
