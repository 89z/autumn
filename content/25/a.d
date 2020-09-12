import std.conv, std.stdio;

void main() {
   auto s = "2.9";
   auto n = s.to!(float);
   writeln(n == 2.9f);
}
