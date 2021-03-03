import file = std.file;
import io = std.stdio;

void main() {
   auto t = file.timeLastModified("a.d");
   io.writeln(t);
}
