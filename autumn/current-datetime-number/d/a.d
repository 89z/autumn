import
   io = std.stdio,
   std.datetime: Clock;

void main() {
   auto o = Clock.currTime;
   auto n = o.toUnixTime;
   io.writeln(n);
}
