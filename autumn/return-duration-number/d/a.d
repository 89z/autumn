import io = std.stdio;
import std.datetime: Clock, SysTime;

void main() {
   auto t = SysTime.fromISOExtString("2020-12-31T01:02:31");
   auto n = (Clock.currTime - t).total!"seconds";
   io.writeln(n);
}
