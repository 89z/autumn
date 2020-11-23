import std.conv: text;
import std.stdio: writeln;

void main() {
   auto n = 100;
   auto s = text(n);
   writeln(s == "100");
}
