import conv = std.conv;
import io = std.stdio;

void main() {
   auto s = conv.text(100);
   io.writeln(s == "100");
}
