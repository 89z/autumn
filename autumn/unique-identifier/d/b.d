import conv = std.conv;
import io = std.stdio;

void main() {
   // example 1
   auto n = 0xFFFF_FFFF;
   auto s = conv.to!string(n, 36);
   // example 2
   auto s2 = "1Z141Z3";
   auto n2 = conv.to!uint(s, 36);
   // print
   io.writeln(s2 == s && n2 == n);
}
