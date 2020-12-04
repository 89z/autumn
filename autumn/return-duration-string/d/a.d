import
   io = std.stdio,
   core.time: MonoTime;

void main() {
   auto o = MonoTime.currTime;
   io.writeln(o);
}
