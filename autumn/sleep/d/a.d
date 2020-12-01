import
   core.thread,
   std.stdio;

void main() {
   auto o = seconds(5);
   writeln("sleep");
   Thread.sleep(o);
}
