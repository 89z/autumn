import
   io = std.stdio,
   time = std.datetime.stopwatch,
   core.thread: Thread;

void main() {
   auto dur_o = time.msecs(10);
   auto sw_o = time.StopWatch();
   sw_o.start;
   int sec_n, mil_n;
   while (true) {
      Thread.sleep(dur_o);
      sw_o.peek.split!("seconds", "msecs")(sec_n, mil_n);
      io.writef("\r%d s %03d ms", sec_n, mil_n);
   }
}
