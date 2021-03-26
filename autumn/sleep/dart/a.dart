import 'dart:io';

void main() {
   var d = Duration(milliseconds: 500);
   while (true) {
      print('sleep');
      sleep(d);
   }
}
