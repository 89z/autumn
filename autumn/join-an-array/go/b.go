package main

func main() {
   a := []byte{'w', 'e', 's', 't'}
   s := string(a)
   println(s == "west")
}
