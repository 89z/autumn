import std.path, std.stdio;
void main() {
   auto s1 = "index.md";
   auto s2 = s1.extension;
   writeln(s2 == ".md");
}
