import std.datetime: SysTime;
import std.file: setTimes;

void main() {
   auto o = SysTime.fromUnixTime(400_000_000);
   setTimes("a.d", o, o);
}
