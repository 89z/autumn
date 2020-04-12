package main
import "strings"
func u_group(s_in string) string {
   var n_neg = strings.LastIndex("-" + s_in, "-")
   var n_dot = strings.Index(s_in + ".", ".")
   var s_out string
   for n_pos, s_pos := range s_in {
      if n_neg < n_pos && n_pos < n_dot && n_pos % 3 == n_dot % 3 {
         s_out += ","
      }
      s_out += string(s_pos)
   }
   return s_out
}
func main() {
   a1 := []string{"123", "1234", "-123", "-1234", "123.4", "1234.5"}
   for _, s1 := range a1 {
      println(u_group(s1))
   }
}
