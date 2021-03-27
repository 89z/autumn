import 'dart:io';

void main() {
   var f = new File('a.txt');
   f.deleteSync();
}
