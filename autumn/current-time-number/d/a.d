import
   std.datetime,
   std.stdio;

void main() {
   auto o = Clock.currTime;
   auto n = o.toUnixTime;
   n.writeln;
}
