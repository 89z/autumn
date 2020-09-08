package main
import "fmt"
type flip map[string]int

func f(m_acc flip, s_cur string, n_idx int) flip {
   m_acc[s_cur] = n_idx
   return m_acc
}

func main() {
   a := []string{"May", "June"}
   m := flip{}
   for n, s := range a {
      m = f(m, s, n)
   }
   fmt.Println(m)
}
