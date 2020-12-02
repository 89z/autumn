import std.algorithm: startsWith;
import std.stdio: writeln;

void main() {
   auto s = "March";
   auto b = s.startsWith("Ma");
   b.writeln;
}
