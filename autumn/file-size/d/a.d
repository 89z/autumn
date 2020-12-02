import
   file = std.file,
   io = std.stdio;

void main() {
   auto n = file.getSize("a.d");
   io.writeln(n);
}
