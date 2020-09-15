package main

func reduce(s_acc, s_cur string) string {
   return s_acc + s_cur
}

func main() {
   a_src := []string{"May", "June"}
   s_acc := ""
   for n_idx := range a_src {
      s_cur := a_src[n_idx]
      s_acc = reduce(s_acc, s_cur)
   }
   println(s_acc)
}
