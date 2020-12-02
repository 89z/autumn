import
   conv = std.conv,
   io = std.stdio;

void main() {
   auto s = conv.to!string(100);
   io.writeln(s == "100");
}
