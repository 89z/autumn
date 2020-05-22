package main
import o1 "golang.org/x/sys/windows/registry"
func main() {
   o2, e1 := o1.OpenKey(o1.CURRENT_USER, "Console", o1.SET_VALUE)
   if e1 != nil {
      log.Fatal(e1)
   }
   o2.SetDWordValue("CodePage", 65001)
}
