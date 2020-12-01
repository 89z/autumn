import
   core.thread,
   std.stdio;

void main() {
   auto o = msecs(500);
   while (true) {
      writeln("sleep");
      Thread.sleep(o);
   }
}
