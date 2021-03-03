import 'dart:io';

void main() {
   var f = new File('a.dart');
   var t = f.lastModifiedSync();
   print(t);
}
