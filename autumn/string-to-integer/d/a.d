import
   conv = std.conv,
   io = std.stdio;

void main() {
   auto n = conv.to!int("100");
   io.writeln(n == 100);
}
