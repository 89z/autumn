import file = std.file;
import io = std.stdio;

void main() {
   auto n = file.getSize("a.d");
   io.writeln(n);
}
