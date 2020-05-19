package main
import "golang.org/x/sys/windows"
func main() {
   n1, n2, n3 := windows.RtlGetNtVersionNumbers()
   println(n1, n2, n3)
}
