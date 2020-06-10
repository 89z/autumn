import std.datetime, std.stdio;
void main() {
   auto o = Date(2019, 12, 31);
   auto s = toISOString(o);
   writeln(s);
}
