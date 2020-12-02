import core.thread: Thread, msecs;
import std.stdio: writeln;

void main() {
   auto o = msecs(500);
   while (true) {
      writeln("sleep");
      Thread.sleep(o);
   }
}
