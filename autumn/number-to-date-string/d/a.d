import
   fmt = std.format,
   io = std.stdio,
   std.datetime: SysTime;

void main() {
   auto o = SysTime.fromUnixTime(1577858399);
   auto s = fmt.format("%d-%d-%d", o.year, o.month, o.day);
   io.writeln(s == "2019-12-31");
}
