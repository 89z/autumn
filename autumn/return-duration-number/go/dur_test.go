package dur

import (
   "fmt"
   "testing"
)

var tests = []string{
   "2020",
   "2020-12",
   "2020-12-31",
   "2020-12-31T01",
   "2020-12-31T01:02",
   "2020-12-31T01:02:31",
}

func TestDur(t *testing.T) {
   for _, test := range tests {
      {
         n, err := since(test)
         if err != nil {
            t.Fatal(err)
         }
         fmt.Println(n)
      }
      {
         n, err := sub(test)
         if err != nil {
            t.Fatal(err)
         }
         fmt.Println(n)
      }
   }
}
