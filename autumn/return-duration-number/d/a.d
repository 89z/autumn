import io = std.stdio;
import std.datetime: Clock, SysTime;

void main() {
   auto old_o = SysTime.fromISOExtString("2019-12-31T23:59:59");
   auto new_o = Clock.currTime;
   auto n = (new_o - old_o).total!"days";
   io.writeln(n);
}
