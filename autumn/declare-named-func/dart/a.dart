// example 1
num add(num n, num n1) {
   return n + n1;
}
// example 2
num sub(num n, num n2) => n - n2;
// print
void main() {
   var n1 = add(8, 1);
   var n2 = sub(8, 1);
   print(n1 == 9 && n2 == 7);
}
