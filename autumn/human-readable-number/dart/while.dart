String numberFormat(double d) {
   var e = d, f = 0;
   while (e >= 1000) {
      e /= 1000;
      f++;
   }
   return e.toStringAsFixed(3) + ['', ' k', ' M', ' G'][f];
}

void main() {
   var s = numberFormat(9012345678);
   print(s == '9.012 G');
}
