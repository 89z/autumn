package main
import "os"
func main() {
   // example 1
   os.Mkdir("Sunday", os.ModePerm)
   // exmaple 2
   os.MkdirAll("Monday/Tuesday", os.ModePerm)
}
