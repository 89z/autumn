import io = std.stdio;
import std.datetime: Clock, SysTime;

long sinceHours(string s) {
   auto t = SysTime.fromISOExtString(s);
   return (Clock.currTime - t).total!"hours";
}

void main() {
   auto n = sinceHours("2020-12-31T01:02:31");
   io.writeln(n);
}
