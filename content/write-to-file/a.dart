import 'dart:io';
main() {
   var o = new File('a.txt');
   o.writeAsStringSync('Sunday\n');
}
