import 'dart:io';

void main() {
   { // example 1
      var d = new Directory('March');
      d.createSync();
   }
   { // example 2
      var d = new Directory('May/June');
      d.createSync(recursive: true);
   }
}
