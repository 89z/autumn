import 'dart:io';

void main() {
   var f = new File('a.txt');
   var s = f.readAsStringSync();
   print(s);
}
