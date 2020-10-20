package main

func HasPrefix(s string, c byte) bool {
   return s[0] == c
}

func main() {
   s := "June"
   b := HasPrefix(s, 'J')
   println(b)
}
