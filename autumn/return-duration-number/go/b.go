package main
import "time"

func TimeHours(value string) (float64, error) {
   layout := time.RFC3339[:len(value)]
   o, e := time.Parse(layout, value)
   if e != nil {
      return 0, e
   }
   return time.Since(o).Hours(), nil
}
