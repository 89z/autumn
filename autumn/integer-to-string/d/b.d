import conv = std.conv;
import io = std.stdio;

void main() {
   auto s = conv.to!string(100);
   io.writeln(s == "100");
}
