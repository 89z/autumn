import io = std.stdio;
import std.datetime: Clock;

void main() {
   auto o = Clock.currTime;
   auto s = o.toISOExtString;
   io.writeln(s);
}
