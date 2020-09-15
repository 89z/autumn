package main

func main() {
   a_src := []string{"May", "June"}
   s_acc := ""
   for n_idx := range a_src {
      s_acc += a_src[n_idx]
   }
   println(s_acc)
}
