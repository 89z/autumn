import file = std.file;
import io = std.stdio;

void main() {
   auto o = file.timeLastModified("a.d");
   io.writeln(o);
}
