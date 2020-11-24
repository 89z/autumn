import
   std.file,
   std.stdio;

void main() {
   auto o = timeLastModified("a.d");
   o.writeln;
}
