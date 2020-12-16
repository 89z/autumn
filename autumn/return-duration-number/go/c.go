package main
import "time"

func HoursSince(value string) (float64, error) {
   layout := time.RFC3339[:len(value)]
   o, e := time.Parse(layout, value)
   if e != nil {
      return 0, e
   }
   return time.Now().Sub(o).Hours(), nil
}
