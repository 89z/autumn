import 'dart:io';

void main() {
   var o = File('a.dart');
   var n = o.lengthSync();
   print(n);
}