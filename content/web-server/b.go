package main
import "net/http"
func main() {
   s1 := ":10"
   s2 := "/prefix/"
   o1 := http.Dir(".")
   o2 := http.FileServer(o1)
   o3 := http.StripPrefix(s2, o2)
   http.Handle(s2, o3)
   println("localhost" + s1 + s2)
   http.ListenAndServe(s1, nil)
}
