import file = std.file;
import io = std.stdio;

bool testFile(string s) {
   return file.exists(s) && file.isFile(s);
}

void main() {
   auto b = testFile("a.d");
   io.writeln(b);
}
