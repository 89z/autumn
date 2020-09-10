import std.datetime, std.stdio;
void main() {
   auto s = "2019-12-31T23:59:59";
   auto o = SysTime.fromISOExtString(s);
   auto n = o.toUnixTime;
   n.writeln;
}
