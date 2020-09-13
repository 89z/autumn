import std.datetime, std.stdio, std.string;

void main() {
   auto n = 1577858399;
   auto o = SysTime.fromUnixTime(n);
   auto s = format("%s %s %s %s", o.dayOfWeek, o.month, o.day, o.year);
   writeln(s == "tue dec 31 2019");
}
