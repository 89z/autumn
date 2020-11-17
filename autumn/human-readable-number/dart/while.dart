String numberFormat(double n) {
   int n2 = 0;
   while (n >= 1000) {
      n /= 1000;
      n2++;
   }
   return n.toStringAsFixed(3) + ['', ' k', ' M', ' G'][n2];
}

void main() {
   var s = numberFormat(9012345678);
   print(s == '9.012 G');
}
