import std.process, std.stdio;
void main() {
   auto s = environment.get("BROWSER");
   s.writeln;
}
