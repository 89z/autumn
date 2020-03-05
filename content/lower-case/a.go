package main
import (
   "golang.org/x/text/cases"
   "golang.org/x/text/language"
   "strings"
)
func main() {
   s1 := "SUNDAY"
   // example 1
   s2 := strings.ToLower(s1)
   // example 2
   s3 := cases.Fold().String(s1)
   // example 3
   s4 := cases.Lower(language.English).String(s1)
   // print
   println(s2, s3, s4)
}
