import core.thread: Thread;
import io = std.stdio;
import time = std.datetime.stopwatch;

void main() {
   auto o = time.StopWatch();
   o.start;
   while (true) {
      Thread.sleep(time.msecs(10));
      io.writef(
         "\r%d m %02d s %03d ms",
         o.peek.split.minutes,
         o.peek.split.seconds,
         o.peek.split.msecs
      );
   }
}
