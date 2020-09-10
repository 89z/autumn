package main
import "os"

func main() {
   // example 1
   os.Mkdir("May", os.ModeDir)
   // exmaple 2
   os.MkdirAll("June/July", os.ModeDir)
}
