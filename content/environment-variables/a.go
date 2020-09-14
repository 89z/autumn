package main
import "os"

func main() {
   s := "USERPROFILE"
   // example 1
   s2 := os.Getenv(s)
   // example 2
   s3, b := os.LookupEnv(s)
   // print
   print(s2, "\n", s3, "\n", b, "\n")
}
