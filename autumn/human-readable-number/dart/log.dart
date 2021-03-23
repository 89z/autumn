import 'dart:math';

String numberFormat(double d) {
   var e = log(d) / ln10 ~/ 3;
   var f = d / pow(1000, e);
   return f.toStringAsFixed(3) + ['', ' k', ' M', ' G'][e];
}

void main() {
   var s = numberFormat(9012345678);
   print(s == '9.012 G');
}
