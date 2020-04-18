import "io/ioutil"
import "net/http"

resp, err := http.Get("http://www.google.com")
if err == nil {
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err == nil {
    fmt.Println(string(body))
  }
}
