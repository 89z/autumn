import std.datetime, std.stdio;

void main() {
   auto o = Clock.currTime;
   auto s = o.toISOExtString;
   s.writeln;
}
