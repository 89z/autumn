void main() {
   var s, n = 1;
   if (n > 0) {
      s = '+';
   } else if (n < 0) {
      s = '-';
   } else {
      s = 'zero';
   }
   print(s == '+');
}
