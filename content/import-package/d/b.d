import std.conv: text;
import std.stdio: writeln;

void main() {
   auto n = 11;
   auto s = text(n);
   writeln(s == "11");
}
