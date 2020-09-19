package main
import "os"

func main() {
   s := "USERPROFILE"
   // example 1
   s1 := os.Getenv(s)
   // example 2
   s2, b2 := os.LookupEnv(s)
   // print
   s0 := `C:\Users\Steven`
   println(s1 == s0, s2 == s0, b2)
}
