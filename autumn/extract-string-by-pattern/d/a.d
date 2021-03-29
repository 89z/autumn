import std.array, std.regex, std.stdio;

void main() {
   auto s = "south north";
   auto a = s.matchAll("o..").array;
   writeln(a[0][0] == "out" && a[1][0] == "ort");
}
