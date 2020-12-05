import 'dart:io';

void main() {
   var o = new Stopwatch();
   o.start();
   while (true) {
      sleep(Duration(milliseconds: 10));
      var s = o.elapsed.toString();
      stdout.write(s + '\r');
   }
}
