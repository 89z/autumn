import core.thread, std.datetime, std.stdio;

void main() {
   auto then = Clock.currTime;
   while (true) {
      Thread.sleep(msecs(10));
      auto now = Clock.currTime - then;
      writef(
         "\r%d m %02d s %03d ms",
         now.split.minutes,
         now.split.seconds,
         now.split.msecs
      );
   }
}
