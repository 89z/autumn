import io = std.stdio;
import std.datetime: SysTime;

void main() {
   auto o = SysTime.fromSimpleString("2020-Dec-31 23:59:59");
   io.writeln(o);
}
