void main() {
   num n;
   print('example 1');
   n = 10;
   while (n < 20) {
      print(n);
      n++;
   }
   print('example 2');
   n = 10;
   while (true) {
      if (n > 19) {
         break;
      }
      print(n);
      n++;
   }
}
