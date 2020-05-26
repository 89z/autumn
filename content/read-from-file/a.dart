import 'dart:io';
main() {
   var o = new File('a.txt');
   var s = o.readAsStringSync();
   print(s);
}
