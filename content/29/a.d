import std.datetime, std.stdio;

void main() {
   auto n = 1577858399;
   auto o = SysTime.fromUnixTime(n);
   auto s = o.toISOExtString;
   writeln(s == "2019-12-31T23:59:59");
}
