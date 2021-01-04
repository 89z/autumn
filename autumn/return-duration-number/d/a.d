import io = std.stdio;
import std.datetime: Clock, SysTime;

void main() {
   auto old_o = SysTime.fromISOExtString("2020-05-04T03:02:01");
   auto new_o = Clock.currTime;
   auto n = (new_o - old_o).total!"days";
   io.writeln(n);
}
