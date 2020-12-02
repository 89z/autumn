import core.time: MonoTime;
import io = std.stdio;

void main() {
   auto o = MonoTime.currTime;
   io.writeln(o);
}
