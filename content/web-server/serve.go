package main
import "net/http"
func main() {
   s1 := ":10"
   o1 := http.Dir(".")
   o2 := http.FileServer(o1)
   println("localhost" + s1)
   http.ListenAndServe(s1, o2)
}
