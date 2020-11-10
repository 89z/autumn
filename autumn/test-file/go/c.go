package main
import "os"

func main() {
   o, e := os.Stat(`C:\Windows\notepad.exe`)
   if os.IsNotExist(e) {
      os.Exit(1)
   }
   s := o.Name()
   println(s == "notepad.exe")
}
