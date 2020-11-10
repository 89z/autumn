import 'dart:io';

void main() {
   var s = 'March\n';
   var o = new File('a.txt');
   o.writeAsStringSync(s);
}
