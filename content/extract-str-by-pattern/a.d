import  std.regex, std.stdio;
void main() {
   auto s = "Wednesday";
   auto a = s.matchAll("e.".regex);
   a.writeln;
}
