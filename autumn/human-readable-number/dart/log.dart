import 'dart:math';

String numberFormat(double n) {
   var n2 = (log(n) / ln10 / 3).toInt();
   var n3 = n / pow(1000, n2);
   return n3.toStringAsFixed(3) + ['', ' k', ' M', ' G'][n2];
}

void main() {
   var s = numberFormat(9012345678);
   print(s == '9.012 G');
}
