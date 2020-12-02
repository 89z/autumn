import
   io = std.stdio,
   time = std.datetime.stopwatch,
   core.thread: Thread;

void main() {
   auto dur_o = time.msecs(10);
   auto sw_o = time.StopWatch();
   sw_o.start;
   while (true) {
      Thread.sleep(dur_o);
      io.write(sw_o.peek.total!"seconds", sw_o.peek.total!"msecs", "\r");
   }
}
