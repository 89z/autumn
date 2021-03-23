import core.thread: Thread;
import io = std.stdio;
import time = std.datetime.stopwatch;

void main() {
   auto t = time.StopWatch();
   t.start;
   while (true) {
      Thread.sleep(time.msecs(10));
      io.writef(
         "\r%d m %02d s %03d ms",
         t.peek.split.minutes,
         t.peek.split.seconds,
         t.peek.split.msecs
      );
   }
}
