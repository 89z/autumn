import core.thread: Thread;
import io = std.stdio;
import time = core.time;

void main() {
   auto t = time.msecs(500);
   while (true) {
      io.writeln("sleep");
      Thread.sleep(t);
   }
}
