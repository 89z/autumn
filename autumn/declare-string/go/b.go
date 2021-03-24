package main

func main() {
   var s string

   // slash
   s = `south\north`
   println(s)

   // quote
   s = `south"north`
   println(s)

   // newline
   s = `south
north`
   println(s)

}
