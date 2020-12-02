import std.path: extension;
import std.stdio: writeln;

void main() {
   auto s = extension("a.d");
   writeln(s == ".d");
}
