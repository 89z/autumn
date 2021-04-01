package main

func main() {
   b := []byte{'w', 'e', 's', 't'}
   s := string(b)
   println(s == "west")
}
