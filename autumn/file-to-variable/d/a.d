import
   file = std.file,
   io = std.stdio;

void main() {
   auto s = file.readText("a.d");
   io.write(s);
}
