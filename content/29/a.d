import std.datetime, std.stdio;

void main() {
   auto n = 366 * 24 * 60 * 60;
   auto o = SysTime.fromUnixTime(n);
   auto s = o.toISOExtString;
   s.writeln;
}
