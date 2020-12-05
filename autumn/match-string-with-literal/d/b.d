import algo = std.algorithm;
import io = std.stdio;

void main() {
   auto s = "March";
   auto b = algo.startsWith(s, "Ma");
   io.writeln(b);
}
