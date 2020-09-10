import std.conv, std.stdio;
void main() {
   auto s = "1.9";
   auto n = s.to!(float);
   writeln(n);
}
