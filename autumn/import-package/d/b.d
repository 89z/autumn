import std.conv: text;
import std.stdio: writeln;

void main() {
   auto s = text(100);
   writeln(s == "100");
}
