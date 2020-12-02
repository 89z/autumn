import
   file = std.file,
   io = std.stdio;

void main() {
   auto o = file.timeLastModified("a.d");
   io.writeln(o);
}
