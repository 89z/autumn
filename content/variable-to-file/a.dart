import 'dart:io';

main() {
   var s = 'May\n';
   var o = new File('a.txt');
   o.writeAsStringSync(s);
}
