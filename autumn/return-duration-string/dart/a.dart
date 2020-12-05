import 'dart:io';

String format(Duration o) {
   var mil_s = (o.inMilliseconds % 1000).toString().padLeft(3, '0');
   var sec_s = (o.inSeconds % 60).toString().padLeft(2, '0');
   return o.inMinutes.toString() + ' m ' + sec_s + ' s ' + mil_s + ' ms';
}

void main() {
   var o = new Stopwatch();
   o.start();
   while (true) {
      sleep(Duration(milliseconds: 10));
      var s = format(o.elapsed);
      stdout.write(s + '\r');
   }
}
