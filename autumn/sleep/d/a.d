import
   io = std.stdio,
   time = core.time,
   core.thread: Thread;

void main() {
   auto o = time.msecs(500);
   while (true) {
      io.writeln("sleep");
      Thread.sleep(o);
   }
}
