import 'dart:io';

void main() {
   // example 1
   var o = new Directory('March');
   o.createSync();
   // example 2
   var o2 = new Directory('May/June');
   o2.createSync(recursive: true);
}
