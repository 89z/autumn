import std.array, std.regex, std.stdio;

void main() {
   auto s = "Sunday Monday";
   auto a = s.matchAll("..n").array;
   writeln(a[0][0] == "Sun" && a[1][0] == "Mon");
}
