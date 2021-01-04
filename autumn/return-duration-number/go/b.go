package main
import "time"

func SinceHours(left string) (float64, error) {
   right := "1970-01-01T00:00:00Z"[len(left):]
   o, e := time.Parse(time.RFC3339, left + right)
   if e != nil {
      return 0, e
   }
   return time.Since(o).Hours(), nil
}
