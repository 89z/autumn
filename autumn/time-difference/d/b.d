import core.thread: Thread;
import std.datetime.stopwatch: StopWatch, msecs;
import std.stdio: write;

void main() {
   auto dur_o = msecs(10);
   auto time_o = StopWatch();
   time_o.start;
   while (true) {
      Thread.sleep(dur_o);
      write(time_o.peek, "\r");
   }
}
