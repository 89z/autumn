package main
import "net/http"
func main() {
   o1 := http.Dir("")
   o2 := http.FileServer(o1)
   http.ListenAndServe(":10", o2)
}
