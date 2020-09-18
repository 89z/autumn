// example A
bool fA(num n, num n2) {
   return n > n2;
}
// example B
bool fB(num n, num n2) => n > n2;
// print
void main() {
   var b = f(9, 8);
   var b2 = f2(9, 8);
   print(b && b2);
}
