import std.datetime, std.stdio;
void main() {
   auto s = "2019-12-31";
   auto o = Date.fromISOExtString(s);
   writeln(o);
}
