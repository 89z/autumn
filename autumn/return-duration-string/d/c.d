import core.thread: Thread;
import io = std.stdio;
import time = std.datetime: Clock;

void main() {
   auto start_o = Clock.currTime;
   while (true) {
      Thread.sleep(time.msecs(10));
      auto o = Clock.currTime - start_o;
      io.writef("\r%d s %03d ms", o.split.seconds, o.split.msecs);
   }
}
