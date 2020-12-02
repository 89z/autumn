import
   conv = std.conv,
   io = std.stdio;

void main() {
   auto n = conv.to!float("9.9");
   io.writeln(n == 9.9f);
}
