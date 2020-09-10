import 'dart:io';
main() {
   var o = new File('index.md');
   var b = o.existsSync();
   print(b);
}
