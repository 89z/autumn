import 'dart:io';
main() {
   // example 1
   var o1 = new Directory('Sunday');
   o1.createSync();
   // example 2
   var o2 = new Directory('Monday/Tuesday');
   o2.createSync(recursive: true);
}
