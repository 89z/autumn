package main
import "os"

func main() {
   s, b := os.LookupEnv("USERPROFILE")
   println(s == `C:\Users\Steven`, b)
}
