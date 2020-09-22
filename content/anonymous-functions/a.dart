void main() {
   // example 1
   var f1 = (num n, num n1) {
      return n > n1;
   };
   // example 2
   var f2 = (num n, num n2) => n > n2;
   // print
   var b1 = f1(9, 8);
   var b2 = f2(9, 8);
   print(b1 && b2);
}
