import
   std.file,
   std.stdio;

void main() {
   auto a = dirEntries(".", SpanMode.shallow);
   writeln(a);
}
