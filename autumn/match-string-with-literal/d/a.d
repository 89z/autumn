import algo = std.algorithm;
import io = std.stdio;

void main() {
   auto s = "March";
   auto b = algo.canFind(s, "ar");
   io.writeln(b);
}
