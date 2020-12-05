import
   io = std.stdio,
   time = core.time,
   core.thread: Thread;

void main() {
   auto start_o = time.MonoTime.currTime;
   while (true) {
      Thread.sleep(time.msecs(10));
      auto o = time.MonoTime.currTime - start_o;
      io.writef(
         "\r%d m %02d s %03d ms",
         o.split.minutes,
         o.split.seconds,
         o.split.msecs
      );
   }
}
