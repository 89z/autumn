import
   file = std.file, io = std.stdio,
   std.file: SpanMode;

void main() {
   auto a = file.dirEntries(".", SpanMode.shallow);
   io.writeln(a);
}
