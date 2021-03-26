import io = std.stdio;
import std.datetime: Clock;

void main() {
   auto t = Clock.currTime;
   auto s = t.toISOExtString;
   io.writeln(s);
}
