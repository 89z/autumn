import
   io = std.stdio,
   std.datetime: SysTime;

void main() {
   auto s = "2019-12-31T23:59:59";
   auto o = SysTime.fromISOExtString(s);
   io.writeln(o);
}
