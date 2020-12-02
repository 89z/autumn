import std.datetime: SysTime;
import std.format: format;
import std.stdio: writeln;

void main() {
   auto o = SysTime.fromUnixTime(1577858399);
   auto s = format("%d-%d-%d", o.year, o.month, o.day);
   writeln(s == "2019-12-31");
}
