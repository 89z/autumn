String numberFormat(double n) {
   int n2 = 0;
   while (n > 1024) {
      n /= 1024;
      n2++;
   }
   return n.toStringAsFixed(3) + ['', ' k', ' M', ' G'][n2];
}

void main() {
   var s = numberFormat(1264);
   print(s == '1.234 k');
}
