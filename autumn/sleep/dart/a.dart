import 'dart:io';

void main() {
   var o = Duration(milliseconds: 500);
   while (true) {
      print('sleep');
      sleep(o);
   }
}
