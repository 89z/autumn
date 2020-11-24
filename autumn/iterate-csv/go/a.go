package main

import (
   "encoding/csv"
   "io"
   "strings"
)

type Scanner struct {
   Reader *csv.Reader
   Head map[string]int
   Row []string
}

func NewScanner(o io.Reader) Scanner {
   csv_o := csv.NewReader(o)
   a, e := csv_o.Read()
   if e != nil {
      return Scanner{}
   }
   m := map[string]int{}
   for n, s := range a {
      m[s] = n
   }
   return Scanner{Reader: csv_o, Head: m}
}

func (o *Scanner) Scan() bool {
   a, e := o.Reader.Read()
   o.Row = a
   return e == nil
}

func (o Scanner) Text(s string) string {
   return o.Row[o.Head[s]]
}

func main() {
   s := `Month,Day
January,Sunday
February,Monday`

   o := NewScanner(strings.NewReader(s))
   for o.Scan() {
      println(o.Text("Month"), o.Text("Day"))
   }
}
