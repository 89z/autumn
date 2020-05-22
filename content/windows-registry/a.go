package main
import o1 "golang.org/x/sys/windows/registry"
func main() {
   o2, e := o1.OpenKey(o1.CURRENT_USER, "Console", o1.SET_VALUE)
   if e != nil {
      log.Fatal(e)
   }
   o2.SetDWordValue("CodePage", 65001)
}
