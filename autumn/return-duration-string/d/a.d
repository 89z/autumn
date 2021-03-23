import core.thread, core.time, std.stdio;

void main() {
   auto then = MonoTime.currTime;
   while (true) {
      Thread.sleep(msecs(10));
      auto now = MonoTime.currTime - then;
      writef(
         "\r%d m %02d s %03d ms",
         now.split.minutes,
         now.split.seconds,
         now.split.msecs
      );
   }
}
