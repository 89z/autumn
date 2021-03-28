import 'dart:io';

void main() {
   { // example 1
      var d = new Directory('north');
      d.createSync();
   }
   { // example 2
      var d = new Directory('west/east');
      d.createSync(recursive: true);
   }
}
