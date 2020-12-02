import
   conv = std.conv,
   io = std.stdio;

void main() {
   auto s = conv.text(100);
   io.writeln(s == "100");
}
