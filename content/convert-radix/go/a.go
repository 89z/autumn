package main
import "math/big"

type Radix struct {
   n_len int
}

func (o_rad Radix) Encode(n_in int64) string {
   return big.NewInt(n_in).Text(o_rad.n_len)
}

func (o_rad Radix) Decode(s_in string) int64 {
   o_big := big.Int{}
   o_big.SetString(s_in, o_rad.n_len)
   return o_big.Int64()
}

func main() {
   n := int64(1577858399)
   o := Radix{62}
   // example 1
   s1 := o.Encode(n)
   n1 := o.Decode(s1)
   // example 2
   s2 := o.Encode(n - 1)
   n2 := o.Decode(s2)
   // print
   println(s1 == "1IMx31", n1 == n, s2 == "1IMx30", n2 == n - 1)
}
