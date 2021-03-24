import io = std.stdio;
import std.path;

void main() {
   auto s = buildPath("south", "north");
   io.writeln(s == `south\north`);
}
