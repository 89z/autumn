import
   std.datetime,
   std.stdio,
   std.string;

void main() {
   auto o = SysTime.fromUnixTime(1577858399);
   auto s = format("%d-%d-%d", o.year, o.month, o.day);
   writeln(s == "2019-12-31");
}
