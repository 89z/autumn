import 'dart:io';

void main() {
   var f = new File('index.md');
   var b = f.existsSync();
   print(b);
}
