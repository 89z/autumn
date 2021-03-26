import io = std.stdio;
import std.datetime: SysTime;

void main() {
   auto t = SysTime.fromSimpleString("2020-Dec-31 01:02:31");
   io.writeln(t);
}
