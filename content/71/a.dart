import 'dart:io';

void main() {
   var o = new File('a.txt');
   o.copySync('b.txt');
}
