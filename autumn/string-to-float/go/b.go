package main
import "fmt"

func FloatVal(s string) (float64, error) {
   var value float64
   count, e := fmt.Sscan(s, &value)
   if e != nil {
      return 0, fmt.Errorf("%v %v", count, e)
   }
   return value, nil
}
