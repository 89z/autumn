void main() {
   print('example 1');
   var n1 = 10;
   while (n1 < 20) {
      print(n1);
      n1++;
   }

   print('example 2');
   var n2 = 10;
   while (true) {
      if (n2 > 19) {
         break;
      }
      print(n2);
      n2++;
   }
}
