import 'dart:io';
main() {
   var o = new Directory('.');
   var a = o.listSync();
   print(a);
}
