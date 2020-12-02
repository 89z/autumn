import core.time: MonoTime;
import std.stdio: writeln;

void main() {
   auto o = MonoTime.currTime;
   o.writeln;
}
