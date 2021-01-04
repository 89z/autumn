import io = std.stdio;
import std.datetime: SysTime;

void main() {
   auto o = SysTime.fromSimpleString("2020-May-04 03:02:01");
   io.writeln(o);
}
