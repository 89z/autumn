import std.datetime, std.stdio;

void main() {
   auto n = 0x5555_5555;
   auto t = SysTime.fromUnixTime(n);
   writeln(t);
}
