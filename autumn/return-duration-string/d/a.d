import core.thread: Thread;
import io = std.stdio;
import time = core.time: MonoTime;

void main() {
   auto from = MonoTime.currTime;
   while (true) {
      Thread.sleep(time.msecs(10));
      auto to = MonoTime.currTime - from;
      io.writef(
         "\r%d m %02d s %03d ms",
         to.split.minutes,
         to.split.seconds,
         to.split.msecs
      );
   }
}
