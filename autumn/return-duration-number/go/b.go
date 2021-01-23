package main
import "time"

func sinceHours(left string) (float64, error) {
   right := "1970-01-01T00:00:00Z"[len(left):]
   t, err := time.Parse(time.RFC3339, left + right)
   if err != nil {
      return 0, err
   }
   return time.Since(t).Hours(), nil
}
