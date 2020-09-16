// example 1
bool f(num n, num n2) {
   return n > n2;
}
// example 2
bool f2(num n, num n2) => n > n2;
// example 3
var f3 = (num n, num n2) => n > n2;
// print
void main() {
   var b = f(9, 8);
   var b2 = f2(9, 8);
   var b3 = f3(9, 8);
   print([b, b2, b3]);
}
