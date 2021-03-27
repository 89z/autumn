import 'dart:io';

void main() {
   var d = new Directory('.');
   var a = d.listSync();
   print(a);
}
