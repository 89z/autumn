package main
import wr "golang.org/x/sys/windows/registry"
func main() {
   k1, _ := wr.OpenKey(wr.CURRENT_USER, "Console", wr.SET_VALUE)
   k1.SetDWordValue("CodePage", 65001)
}
