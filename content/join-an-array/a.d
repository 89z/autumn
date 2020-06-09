import std.array, std.stdio;
void main() {
   auto a = ["Sun", "Mon"];
   auto s = a.join;
   writeln(s == "SunMon");
}
