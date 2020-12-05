import file = std.file;
import std.datetime: SysTime;

void main() {
   auto o = SysTime.fromUnixTime(400_000_000);
   file.setTimes("a.d", o, o);
}
