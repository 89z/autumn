String numberFormat(double n) {
   var n2 = n, n3 = 0;
   while (n2 >= 1000) {
      n2 /= 1000;
      n3++;
   }
   return n2.toStringAsFixed(3) + ['', ' k', ' M', ' G'][n3];
}

void main() {
   var s = numberFormat(9012345678);
   print(s == '9.012 G');
}
