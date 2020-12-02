import std.algorithm: canFind;
import std.stdio: writeln;

void main() {
   auto s = "March";
   auto b = s.canFind("ar");
   b.writeln;
}
