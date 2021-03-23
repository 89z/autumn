import 'dart:io';

String format(Duration d) {
   var mil = (d.inMilliseconds % 1000).toString().padLeft(3, '0');
   var sec = (d.inSeconds % 60).toString().padLeft(2, '0');
   return d.inMinutes.toString() + ' m ' + sec + ' s ' + mil + ' ms';
}

void main() {
   var then = new Stopwatch();
   then.start();
   while (true) {
      sleep(Duration(milliseconds: 10));
      var now = format(then.elapsed);
      stdout.write(now + '\r');
   }
}
