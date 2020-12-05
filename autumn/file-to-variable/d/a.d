import file = std.file;
import io = std.stdio;

void main() {
   auto s = file.readText("a.d");
   io.write(s);
}
