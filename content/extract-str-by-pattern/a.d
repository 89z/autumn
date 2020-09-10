import  std.regex, std.stdio;

void main() {
   auto s = "January";
   auto a = s.matchAll("a.".regex);
   a.writeln;
}
