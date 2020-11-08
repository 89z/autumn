import 'dart:io';

void main() {
   var o = new Directory('.');
   var a = o.listSync();
   print(a);
}
