import file = std.file: SpanMode;
import io = std.stdio;

void main() {
   auto a = file.dirEntries(".", SpanMode.shallow);
   io.writeln(a);
}
