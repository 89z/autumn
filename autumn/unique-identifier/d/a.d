import conv = std.conv;
import io = std.stdio;

string idEncode(int year) {
   conv.to!string(n, 36);
}

void main() {
   auto s = idEncode(2021);
   io.writeln(s);
}
