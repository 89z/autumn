import core.thread: Thread;
import io = std.stdio;
import time = core.time: MonoTime;

void main() {
   auto start_o = MonoTime.currTime;
   while (true) {
      Thread.sleep(time.msecs(10));
      auto o = MonoTime.currTime - start_o;
      io.writef(
         "\r%d m %02d s %03d ms",
         o.split.minutes,
         o.split.seconds,
         o.split.msecs
      );
   }
}
