package main
import "net/http"

func main() {
   r, e := http.NewRequest("GET", "https://stackoverflow.com", nil)
   if e != nil {
      panic(e)
   }
   r.Header.Set("Cookie", "west=left")
}
