import
   io = std.stdio,
   std.datetime: SysTime;

void main() {
   auto o = SysTime.fromSimpleString("2019-Dec-31 23:59:59");
   io.writeln(o);
}
