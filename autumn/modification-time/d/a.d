import
   std.datetime,
   std.file;

void main() {
   auto o = SysTime.fromUnixTime(400_000_000);
   setTimes("a.d", o, o);
}
