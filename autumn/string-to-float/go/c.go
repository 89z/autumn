package main
import "strconv"

func FloatVal(s string) (float64, error) {
   return strconv.ParseFloat(s, 64)
}
