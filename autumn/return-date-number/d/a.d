import io = std.stdio;
import std.datetime: Clock;

void main() {
   auto t = Clock.currTime;
   auto n = t.fracSecs.total!"msecs";
   io.writeln(n);
}
