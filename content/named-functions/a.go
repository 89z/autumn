package main
// example A
func fA(nY int, nZ int) bool {
   return nY > nZ
}
// example B
func fB(nY, nZ int) bool {
   return nY > nZ
}
// print
func main() {
   bA := fA(9, 8)
   bB := fB(9, 8)
   println(bA, bB)
}
