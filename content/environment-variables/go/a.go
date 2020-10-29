package main
import "os"

func main() {
   s := os.Getenv("USERPROFILE")
   println(s == `C:\Users\Steven`)
}
