import io = std.stdio;
import path = std.path;

void main() {
   auto s = path.extension("a.d");
   io.writeln(s == ".d");
}
