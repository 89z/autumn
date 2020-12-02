import core.thread: Thread;
import io = std.stdio;
import time = std.datetime.stopwatch;

void main() {
   auto dur_o = time.msecs(10);
   auto time_o = time.StopWatch();
   time_o.start;
   while (true) {
      Thread.sleep(dur_o);
      io.write(time_o.peek, "\r");
   }
}
