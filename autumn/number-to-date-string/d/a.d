import
   std.datetime,
   std.stdio,
   std.string;

void main() {
   auto o = SysTime.fromUnixTime(1577858399);
   auto s = format("%s %s %s %s", o.dayOfWeek, o.month, o.day, o.year);
   writeln(s == "tue dec 31 2019");
}
