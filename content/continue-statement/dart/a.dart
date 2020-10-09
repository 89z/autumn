void main() {
   for (var n = 0; n < 10; n++) {
      if (n % 3 == 0) {
         continue;
      }
      print(n);
   }
}
