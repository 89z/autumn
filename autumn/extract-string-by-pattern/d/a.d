import std.range, std.regex, std.stdio;

void main() {
   auto o = regex("..n");
   auto a = matchAll("Sunday Monday", o).array;
   writeln(a[0][0] == "Sun" && a[1][0] == "Mon");
}
