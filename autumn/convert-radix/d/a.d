import conv = std.conv;
import io = std.stdio;

void main() {
   // example 1
   auto n = 1609480799;
   auto s1 = conv.to!string(n, 36);
   // example 2
   auto s = "Q3EZBZ";
   auto n2 = conv.to!int(s, 36);
   // print
   io.writeln(s1 == s && n2 == n);
}
