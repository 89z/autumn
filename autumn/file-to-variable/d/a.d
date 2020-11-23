import
   std.file,
   std.stdio;

void main() {
   auto s = readText("a.d");
   s.write;
}
