import file = std.file;
import io = std.stdio;

bool testDir(string s) {
   return file.exists(s) && file.isDir(s);
}

void main() {
   auto b = testDir(`C:\Users`);
   io.writeln(b);
}
