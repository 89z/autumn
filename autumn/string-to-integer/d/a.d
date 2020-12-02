import std.conv: to;
import std.stdio: writeln;

void main() {
   auto s = "100";
   auto n = s.to!int;
   writeln(n == 100);
}
