package main

func main() {
   n := 1 << 32
   println(n == 0xFFFF_FFFF + 1)
}
