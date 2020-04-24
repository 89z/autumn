package main
import o1 "golang.org/x/sys/windows/registry"
func main() {
   o2, _ := o1.OpenKey(o1.CURRENT_USER, "Console", o1.SET_VALUE)
   o2.SetDWordValue("CodePage", 65001)
}
