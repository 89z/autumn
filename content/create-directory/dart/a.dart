import 'dart:io';

void main() {
   // example 1
   var o = new Directory('May');
   o.createSync();
   // example 2
   var o2 = new Directory('June/July');
   o2.createSync(recursive: true);
}
