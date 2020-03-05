package main
import "net/http"
func main() {
   t1 := http.Dir(".")
   f1 := http.FileServer(t1)
   http.ListenAndServe(":10", f1)
}
