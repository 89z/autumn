import io = std.stdio;
import std.datetime: SysTime;

void main() {
   auto s = "2020-05-04T03:02:01";
   auto t = SysTime.fromISOExtString(s);
   io.writeln(t);
}
